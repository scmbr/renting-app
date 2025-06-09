import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "@/shared/api/axios";
import FeatureCard from "@/entities/feature-card/FeatureCard";
import { motion } from "framer-motion";
import { Carousel } from "react-responsive-carousel";
import "react-responsive-carousel/lib/styles/carousel.min.css";
import styles from "./EditAdvertForm.module.css";

const EditAdvertForm = ({ apartment, advertId, initialData }) => {
  const [form, setForm] = useState({
    apartment_id: apartment ? apartment.id : null,
    title: "",
    rent: "",
    deposit: "",
    rental_type: "monthly",
    pets: false,
    babies: false,
    smoking: false,
    internet: false,
    washing_machine: false,
    tv: false,
    conditioner: false,
    dishwasher: false,
    concierge: false,
  });
  const featureIcons = {
    pets: { label: "Можно с животными", src: "/icons/pets.svg" },
    babies: { label: "Можно с детьми", src: "/icons/baby.svg" },
    smoking: { label: "Можно курить", src: "/icons/smoking.svg" },
    internet: { label: "Интернет", src: "/icons/internet.svg" },
    washing_machine: { label: "Стиралка", src: "/icons/washing_machine.svg" },
    tv: { label: "Телевизор", src: "/icons/tv.svg" },
    conditioner: { label: "Кондиционер", src: "/icons/conditioner.svg" },
    dishwasher: { label: "Посудомойка", src: "/icons/dishwasher.svg" },
    concierge: { label: "Консьерж", src: "/icons/concierge.svg" },
  };
  const [photos, setPhotos] = useState([]);
  const [loadingPhotos, setLoadingPhotos] = useState(false);

  const navigate = useNavigate();

  useEffect(() => {
    if (apartment && initialData) {
      setForm({
        ...initialData,
        apartment_id: apartment.id,
        rent: String(initialData.rent),
        deposit: String(initialData.deposit),
      });
      fetchPhotos(apartment.id);
    } else if (apartment) {
      fetchPhotos(apartment.id);
    }
  }, [initialData, apartment]);

  const fetchPhotos = async (apartmentId) => {
    setLoadingPhotos(true);
    try {
      const res = await api.get(`/apartment/${apartmentId}/photos`);
      const sortedPhotos = [...res.data].sort(
        (a, b) => b.is_cover - a.is_cover
      );
      setPhotos(sortedPhotos);
    } catch (e) {
      console.error("Ошибка загрузки фото", e);
      setPhotos([]);
    } finally {
      setLoadingPhotos(false);
    }
  };

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const normalizedForm = {
      ...form,
      rent: Number(form.rent),
      deposit: Number(form.deposit),
    };

    try {
      await api.put(`/my/advert/${advertId}`, normalizedForm);
      navigate("/my/advert");
    } catch (err) {
      console.error("Ошибка при обновлении объявления", err);
    }
  };

  if (!apartment) {
    return <p>Загрузка квартиры...</p>;
  }
  return (
    <motion.form
      onSubmit={handleSubmit}
      className={styles.form}
      animate={{ scale: 1.02 }}
      transition={{ duration: 0.3 }}
    >
      <div className={styles.inputs}>
        <div className={styles.inputRow}>
          <label htmlFor="title">Заголовок:</label>
          <input
            id="title"
            name="title"
            value={form.title}
            onChange={handleChange}
            required
            className={styles.inputField}
          />
        </div>

        <div className={styles.inputRow}>
          <label htmlFor="rent">Аренда:</label>
          <input
            type="number"
            id="rent"
            name="rent"
            value={form.rent}
            onChange={handleChange}
            required
            className={styles.inputField}
          />
        </div>

        <div className={styles.inputRow}>
          <label htmlFor="deposit">Депозит:</label>
          <input
            type="number"
            id="deposit"
            name="deposit"
            value={form.deposit}
            onChange={handleChange}
            required
            className={styles.inputField}
          />
        </div>

        <div className={styles.inputRow}>
          <label htmlFor="rental_type">Тип аренды:</label>
          <select
            id="rental_type"
            name="rental_type"
            value={form.rental_type}
            onChange={handleChange}
            className={styles.selectField}
          >
            <option value="monthly">Помесячно</option>
            <option value="daily">Посуточно</option>
          </select>
        </div>
      </div>

      <h2 className={styles.formTitle}>Выбранная квартира</h2>
      <div className={styles.apartmentDetails}>
        {loadingPhotos ? (
          <p>Загрузка фото...</p>
        ) : photos.length > 0 ? (
          <Carousel
            showThumbs={false}
            showStatus={false}
            infiniteLoop
            className={styles.carousel}
          >
            {photos.map((photo) => (
              <div key={photo.id}>
                <img
                  src={photo.url}
                  alt="Фото квартиры"
                  className={styles.photo}
                />
              </div>
            ))}
          </Carousel>
        ) : (
          <div className={styles.noPhotoWrapper}>
            <img
              src="/images/no-photo.png"
              alt="Нет фото"
              className={styles.noPhoto}
            />
          </div>
        )}
        <div className={styles.apartmentInfo}>
          <p>
            Адрес: {apartment.city} {apartment.street}, {apartment.building}
          </p>
          <p>Год постройки: {apartment.construction_year}</p>
          <p>Тип: {apartment.construction_type}</p>
          <hr />
          <p>Комнат: {apartment.rooms}</p>
          <p>Площадь: {apartment.area} м²</p>
          <p>Этаж: {apartment.floor}</p>
          <p>Санузел: {apartment.bathroom_type}</p>
          <p>Ремонт: {apartment.remont}</p>
        </div>
      </div>

      <h2 className={styles.formTitle}>Удобства и разрешения</h2>
      <div className={styles.featuresGrid}>
        {Object.entries(featureIcons).map(([key, { label, src }]) => (
          <FeatureCard
            key={key}
            label={label}
            icon={
              <img src={src} alt={label} style={{ width: 24, height: 24 }} />
            }
            selected={form[key]}
            onClick={() => setForm((prev) => ({ ...prev, [key]: !prev[key] }))}
          />
        ))}
      </div>

      <button type="submit" className={styles.submitButton}>
        Сохранить
      </button>
    </motion.form>
  );
};

export default EditAdvertForm;
