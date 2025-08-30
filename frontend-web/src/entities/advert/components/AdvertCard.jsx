import { Link } from "react-router-dom";
import { useState } from "react";
import { Carousel } from "react-responsive-carousel";
import "react-responsive-carousel/lib/styles/carousel.min.css";
import styles from "./AdvertCard.module.css";
import { addToFavorites, removeFromFavorites } from "@/entities/favorites/api";

const AdvertCard = ({ advert, onRemoveFavorite }) => {
  if (!advert || !advert.apartment)
    return <div>Данные объявления недоступны</div>;

  const { id, title, rent, apartment, is_favorite } = advert;
  const [isFavorite, setIsFavorite] = useState(is_favorite || false);

  const handleFavoriteClick = async (e) => {
    e.preventDefault();
    try {
      if (isFavorite) {
        await removeFromFavorites(id);
        setIsFavorite(false);
        if (onRemoveFavorite) onRemoveFavorite(id);
      } else {
        await addToFavorites(id);
        setIsFavorite(true);
      }
    } catch (error) {
      console.error("Ошибка избранного:", error);
    }
  };

  const photos = apartment.apartment_photos || [];

  return (
    <div className={styles.cardWrapper}>
      <div className={styles.advertCard}>
        {photos.length > 0 ? (
          <Carousel
            showThumbs={false}
            infiniteLoop
            showStatus={false}
            swipeable
            emulateTouch
            className={styles.сarousel}
            lazyLoad
          >
            {photos.map((photo, index) => (
              <div
                key={photo.id || index}
                className={styles.slideWrapper}
                style={{
                  "--bg-image": `url(${import.meta.env.VITE_BACKEND_URL}${
                    photo.url
                  })`,
                }}
              >
                <img
                  src={`${import.meta.env.VITE_BACKEND_URL}${photo.url}`}
                  alt={`Фото квартиры ${title}`}
                  className={styles.carouselImg}
                />
              </div>
            ))}
          </Carousel>
        ) : (
          <img
            src="/images/no-photo.png"
            alt="Нет фото"
            className={styles.coverPhoto}
            loading="lazy"
          />
        )}

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

      <Link to={`/advert/${id}`} className={styles.advertLink}>
        <div className={styles.advertInfo}>
          <h3 className={styles.title}>{rent.toLocaleString()} ₽/мес</h3>
          <p>
            Этаж: {apartment.floor} | Комнат: {apartment.rooms} | Площадь:{" "}
            {apartment.area}
          </p>
          <p>
            {apartment.street}, {apartment.building}
          </p>
        </div>
      </Link>
    </div>
  );
};

export default AdvertCard;
