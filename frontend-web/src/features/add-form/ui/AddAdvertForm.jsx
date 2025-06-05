import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "@/shared/api/axios";
import ApartmentCard from "@/entities/apartment/ui/ApartmentCard";
import styles from "./AddAdvertForm.module.css";

const AddAdvertForm = ({ apartments = [] }) => {
  const [form, setForm] = useState({
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
  });
  const checkboxLabels = {
    pets: "Животные",
    babies: "Дети",
    smoking: "Курение",
    internet: "Интернет",
    washing_machine: "Стиральная машина",
    tv: "Телевизор",
    conditioner: "Кондиционер",
    dishwasher: "Посудомойка",
    concierge: "Консьерж",
  };

  useEffect(() => {
    if (apartments && apartments.length > 0) {
      setForm((prev) => ({
        ...prev,
        apartment_id: apartments[0].id,
      }));
    }
  }, [apartments]);

  const navigate = useNavigate();

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
      .then(() => navigate("/my/advert"))
      .catch((err) => console.error("Ошибка при создании объявления", err));
  };

  return (
    <form onSubmit={handleSubmit} className={styles.form}>
      <h2 className={styles.formTitle}>Выберите квартиру:</h2>
      <div className={styles.apartmentsContainer}>
        {(apartments ?? []).map((a) => (
          <ApartmentCard
            key={a.id}
            apartment={a}
            isSelected={form.apartment_id === a.id}
            onSelect={(id) =>
              setForm((prev) => ({ ...prev, apartment_id: id }))
            }
          />
        ))}
        <div
          className={styles.apartmentCard}
          onClick={() => navigate("/apartment/add")}
        >
          <div className={styles.addCardContent}>
            <span className={styles.plusSign}>＋</span>
          </div>
        </div>
      </div>

      <label className={styles.inputLabel}>
        Заголовок:
        <input
          name="title"
          value={form.title}
          onChange={handleChange}
          required
          className={styles.inputField}
        />
      </label>

      <label className={styles.inputLabel}>
        Аренда:
        <input
          type="number"
          name="rent"
          value={form.rent}
          onChange={handleChange}
          required
          className={styles.inputField}
        />
      </label>

      <label className={styles.inputLabel}>
        Депозит:
        <input
          type="number"
          name="deposit"
          value={form.deposit}
          onChange={handleChange}
          required
          className={styles.inputField}
        />
      </label>

      <label className={styles.inputLabel}>
        Тип аренды:
        <select
          name="rental_type"
          value={form.rental_type}
          onChange={handleChange}
          className={styles.selectField}
        >
          <option value="monthly">Помесячно</option>
          <option value="daily">Посуточно</option>
        </select>
      </label>

      <div className={styles.checkboxGroup}>
        {Object.entries(checkboxLabels).map(([key, label]) => (
          <label key={key} className={styles.checkboxLabel}>
            <input
              type="checkbox"
              name={key}
              checked={form[key]}
              onChange={handleChange}
            />
            {label}
          </label>
        ))}
      </div>

      <button type="submit" className={styles.submitButton}>
        Добавить
      </button>
    </form>
  );
};

export default AddAdvertForm;
