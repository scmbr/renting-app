import { Link } from "react-router-dom";
import { useEffect, useState } from "react";
import styles from "./AdvertCard.module.css";
import api from "@/shared/api/axios";
import { addToFavorites, removeFromFavorites } from "@/entities/favorites/api";

const AdvertCard = ({ advert, onRemoveFavorite }) => {
  if (!advert || !advert.apartment) {
    return <div>Данные объявления недоступны</div>;
  }

  const { id, title, rent, apartment } = advert;
  const [coverUrl, setCoverUrl] = useState(null);
  const [isFavorite, setIsFavorite] = useState(false);

  useEffect(() => {
    const fetchPhotos = async () => {
      try {
        const response = await api.get(`/apartment/${apartment.id}/photos`);
        const cover = response.data.find((photo) => photo.is_cover);
        if (cover) setCoverUrl(cover.url);
      } catch (err) {
        console.error("Ошибка при загрузке фото:", err);
      }
    };

    const checkFavorite = async () => {
      const token = localStorage.getItem("accessToken");
      if (!token) return; // Не авторизован — не проверяем избранное

      try {
        const response = await api.get(`/favorites/${id}/check`);
        setIsFavorite(response.data?.is_favorite === true);
      } catch (err) {
        console.error("Ошибка при проверке избранного:", err);
      }
    };

    fetchPhotos();
    checkFavorite();
  }, [apartment.id, id]);

  const handleFavoriteClick = async (e) => {
    e.preventDefault();
    try {
      if (isFavorite) {
        await removeFromFavorites(id);
        setIsFavorite(false);
        if (onRemoveFavorite) {
          onRemoveFavorite(id);
        }
      } else {
        await addToFavorites(id);
        setIsFavorite(true);
      }
    } catch (error) {
      console.error("Ошибка избранного:", error);
    }
  };

  return (
    <div className={styles.cardWrapper}>
      <Link to={`/advert/${id}`} className={styles.advertLink}>
        <div className={styles.advertCard}>
          <img
            src={coverUrl || "/images/no-photo.png"}
            alt="Фото квартиры"
            className={styles.coverPhoto}
          />

          <h3 className="text-lg font-bold">{title}</h3>
          <p>
            {apartment.city}, {apartment.street} {apartment.house}
          </p>
          <p>
            Этаж: {apartment.floor} | Комнат: {apartment.rooms}
          </p>
          <p className="font-semibold">{rent.toLocaleString()} ₽/мес</p>
          <button className={styles.favButton} onClick={handleFavoriteClick}>
            <img
              src={
                isFavorite
                  ? "/icons/favourites-filled.svg"
                  : "/icons/favourites.svg"
              }
              alt={isFavorite ? "В избранном" : "Добавить в избранное"}
            />
          </button>
        </div>
      </Link>
    </div>
  );
};

export default AdvertCard;
