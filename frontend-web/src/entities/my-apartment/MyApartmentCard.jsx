import { useEffect, useState } from "react";
import styles from "./MyApartmentCard.module.css";
import api from "@/shared/api/axios";

const MyApartmentCard = ({ apartment, onEdit, onDelete }) => {
  const coverUrl =
    apartment.apartment_photos && apartment.apartment_photos.length > 0
      ? apartment.apartment_photos.find((photo) => photo.is_cover)?.url ||
        "/images/no-photo.png"
      : "/images/no-photo.png";

  return (
    <div className={styles.card}>
      <div className={styles.actions}>
        <button
          className={styles.actionBtn}
          onClick={() => onEdit && onEdit(apartment.id)}
          aria-label="Редактировать квартиру"
        >
          <img
            src="/icons/edit.svg"
            alt="Редактировать"
            className={styles.icon}
          />
        </button>
        <button
          className={styles.actionBtn}
          onClick={() => onDelete && onDelete(apartment.id)}
          aria-label="Удалить квартиру"
        >
          <img src="/icons/delete.svg" alt="Удалить" className={styles.icon} />
        </button>
      </div>

      <img
        src={
          coverUrl ? `http://localhost:8000${coverUrl}` : "/images/no-photo.png"
        }
        alt="Обложка квартиры"
        className={styles.cover}
        onError={(e) => {
          e.target.onerror = null;
          e.target.src = "/images/no-photo.png";
        }}
      />

      <div className={styles.header}>
        <span>{apartment.city}</span>
        <span>
          {apartment.street} {apartment.building}
        </span>
      </div>

      <div className={styles.info}>
        <p>
          <strong>Комнат:</strong> {apartment.rooms}
        </p>
        <p>
          <strong>Этаж:</strong> {apartment.floor}
        </p>
        <p>
          <strong>Тип дома:</strong> {apartment.construction_type}
        </p>
        <p>
          <strong>Год:</strong> {apartment.construction_year}
        </p>
        <p>
          <strong>Ремонт:</strong> {apartment.remont}
        </p>
        <p>
          <strong>Ванная:</strong> {apartment.bathroom_type}
        </p>
        <p>
          <strong>Лифт:</strong> {apartment.elevator ? "Есть" : "Нет"}
        </p>
        <p>
          <strong>Мусоропровод:</strong>{" "}
          {apartment.garbage_chute ? "Есть" : "Нет"}
        </p>
        <p>
          <strong>Консьерж:</strong> {apartment.concierge ? "Да" : "Нет"}
        </p>
        <p>
          <strong>Статус:</strong> {apartment.status}
        </p>
      </div>
    </div>
  );
};

export default MyApartmentCard;
