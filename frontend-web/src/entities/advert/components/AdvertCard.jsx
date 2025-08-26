import { Link } from "react-router-dom";
import { useState } from "react";
import styles from "./AdvertCard.module.css";
import { addToFavorites, removeFromFavorites } from "@/entities/favorites/api";

const AdvertCard = ({ advert, onRemoveFavorite }) => {
  if (!advert || !advert.apartment) {
    return <div>Данные объявления недоступны</div>;
  }

  const { id, title, rent, apartment, is_favorite } = advert;
  const [isFavorite, setIsFavorite] = useState(is_favorite || false);

  // Берем первую фотку, если есть
  const coverUrl =
    apartment.apartment_photos && apartment.apartment_photos.length > 0
      ? apartment.apartment_photos[0].url
      : null;

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
            src={
              coverUrl
                ? `http://localhost:8000${coverUrl}`
                : "/images/no-photo.png"
            }
            alt="Фото квартиры"
            className={styles.coverPhoto}
          />
          <h3 className="text-lg font-bold">{title}</h3>
          <p>
            {apartment.city}, {apartment.street}, {apartment.building}
          </p>
          <p>
            Этаж: {apartment.floor} | Комнат: {apartment.rooms}
          </p>
          <p className="font-semibold">{rent.toLocaleString()} ₽/мес</p>
          <button className={styles.favButton} onClick={handleFavoriteClick}>
            <img
              src={
                isFavorite
                  ? "/icons/favourites-filled.png"
                  : "/icons/favourites.png"
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
