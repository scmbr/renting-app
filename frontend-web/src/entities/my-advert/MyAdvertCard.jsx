import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./MyAdvertCard.module.css";
import api from "@/shared/api/axios";

const MyAdvertCard = ({ advert, onEdit, onDelete }) => {
  const [coverUrl, setCoverUrl] = useState(null);
  const { id, title, rent, deposit, rental_type, apartment } = advert;
  const navigate = useNavigate();
  const rentalTypeMap = {
    daily: "Посуточно",
    monthly: "Помесячно",
  };
  useEffect(() => {
    const fetchPhotos = async () => {
      try {
        const response = await api.get(`/apartment/${apartment.id}/photos`);
        const cover = response.data.find((photo) => photo.is_cover);
        if (cover) {
          setCoverUrl(cover.url);
        }
      } catch (err) {
        console.error("Ошибка при загрузке фото:", err);
      }
    };

    fetchPhotos();
  }, [apartment.id]);

  const handleCardClick = () => {
    navigate(`/advert/${id}`);
  };

  return (
    <div
      className={styles.card}
      onClick={handleCardClick}
      role="button"
      tabIndex={0}
    >
      <div className={styles.actions}>
        <button
          className={styles.actionBtn}
          onClick={(e) => {
            e.stopPropagation();
            onEdit && onEdit(id);
          }}
          aria-label="Изменить объявление"
        >
          <img src="/icons/edit.svg" alt="Изменить" className={styles.icon} />
        </button>
        <button
          className={styles.actionBtn}
          onClick={(e) => {
            e.stopPropagation();
            onDelete && onDelete(id);
          }}
          aria-label="Удалить объявление"
        >
          <img src="/icons/delete.svg" alt="Удалить" className={styles.icon} />
        </button>
      </div>

      {coverUrl ? (
        <img src={coverUrl} alt="Фото квартиры" className={styles.coverPhoto} />
      ) : (
        <div className={styles.placeholder}>Нет фото</div>
      )}

      <div className={styles.info}>
        <h3 className={styles.title}>{title}</h3>
        <p className={styles.address}>
          {apartment.city}, {apartment.street} {apartment.house}
        </p>
        <p>
          Этаж: {apartment.floor} | Комнат: {apartment.rooms}
        </p>
        <p>
          Тип аренды:{" "}
          <strong>{rentalTypeMap[rental_type] || rental_type}</strong>
        </p>
        <p>
          Депозит: <strong>{deposit.toLocaleString()} ₽</strong>
        </p>
        <p className={styles.price}>
          <strong>{rent.toLocaleString()} ₽</strong>/мес
        </p>
      </div>
    </div>
  );
};

export default MyAdvertCard;
