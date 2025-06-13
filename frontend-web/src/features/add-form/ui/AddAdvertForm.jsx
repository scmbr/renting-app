import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "@/shared/api/axios";
import ApartmentCard from "@/entities/apartment/ui/ApartmentCard";
import FeatureCard from "@/entities/feature-card/FeatureCard";
import { motion, AnimatePresence } from "framer-motion";
import { Carousel } from "react-responsive-carousel";
import "react-responsive-carousel/lib/styles/carousel.min.css";
import styles from "./AddAdvertForm.module.css";
import { useAdvertStore } from "@/stores/useAdvertStore";
const AddAdvertForm = ({ apartments = [] }) => {
  const { filledFormData, setFilledFormData, clearFilledFormData } =
    useAdvertStore();
  console.log(filledFormData);
  const [form, setForm] = useState(
    filledFormData || {
      apartment_id: "",
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
    }
  );
  const navigate = useNavigate();
  const [selectedApartment, setSelectedApartment] = useState(null);
  const [photos, setPhotos] = useState([]);
  const [loadingPhotos, setLoadingPhotos] = useState(false);

  useEffect(() => {
    if (filledFormData && Object.keys(filledFormData).length > 0) {
      setForm(filledFormData);
      const selected = apartments.find(
        (a) => a.id === filledFormData.apartment_id
      );
      if (selected) {
        setSelectedApartment(selected);
        fetchPhotos(selected.id);
      }
    }
  }, [filledFormData, apartments]);

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

  useEffect(() => {
    if (Array.isArray(apartments) && apartments.length > 0) {
      const firstApartment = apartments[0];
      setForm((prev) => ({ ...prev, apartment_id: firstApartment.id }));
      setSelectedApartment(firstApartment);
      fetchPhotos(firstApartment.id);
    }
  }, [apartments]);

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

  const handleSubmit = (e) => {
    e.preventDefault();
    const normalizedForm = {
      ...form,
      rent: Number(form.rent),
      deposit: Number(form.deposit),
    };

    api
      .post("/my/advert", normalizedForm)
      .then(() => {
        clearFilledFormData();
        navigate("/my/advert");
      })
      .catch((err) => console.error("Ошибка при создании объявления", err));
  };

  const handleSelectApartment = (apartment) => {
    setForm((prev) => ({ ...prev, apartment_id: apartment.id }));
    setSelectedApartment(apartment);
    fetchPhotos(apartment.id);
  };
  const navigateToApartment = () => {
    setFilledFormData(form);
    console.log("Заполненные поля: ", filledFormData);
    navigate("/my/apartment/add", { state: { from: "advert" } });
  };
  return (
    <motion.form
      onSubmit={handleSubmit}
      className={styles.form}
      animate={{ scale: form.apartment_id ? 1.02 : 1 }}
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

      <h2 className={styles.formTitle}>Выберите квартиру</h2>

      <div className={styles.apartmentsContainer}>
        {(apartments || []).map((a) => (
          <ApartmentCard
            key={a.id}
            apartment={a}
            isSelected={form.apartment_id === a.id}
            onSelect={() => handleSelectApartment(a)}
          />
        ))}
        <div className={styles.apartmentCard} onClick={navigateToApartment}>
          <div className={styles.addCardContent}>
            <span className={styles.plusSign}>＋</span>
          </div>
        </div>
      </div>

      <AnimatePresence>
        {selectedApartment && (
          <motion.div
            key="apartment-details"
            className={styles.apartmentDetails}
            initial={{ opacity: 0, height: 0 }}
            animate={{ opacity: 1, height: "auto" }}
            exit={{ opacity: 0, height: 0 }}
            transition={{ duration: 0.4 }}
          >
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
                Адрес: {selectedApartment.city} {selectedApartment.street},{" "}
                {selectedApartment.building}
              </p>
              <p>Год постройки: {selectedApartment.construction_year}</p>
              <p>Тип: {selectedApartment.construction_type}</p>
              <hr />
              <p>Комнат: {selectedApartment.rooms}</p>
              <p>Площадь: {selectedApartment.area} м²</p>
              <p>Этаж: {selectedApartment.floor}</p>
              <p>Санузел: {selectedApartment.bathroom_type}</p>
              <p>Ремонт: {selectedApartment.remont}</p>
            </div>
          </motion.div>
        )}
      </AnimatePresence>
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
        Добавить
      </button>
    </motion.form>
  );
};

export default AddAdvertForm;
