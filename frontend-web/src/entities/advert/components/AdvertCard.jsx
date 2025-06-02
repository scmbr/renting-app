import { Link } from "react-router-dom";
import { useEffect, useState } from "react";
import styles from "./AdvertCard.module.css";
import api from "@/shared/api/axios"; // предполагаю, у тебя есть такой инстанс

const AdvertCard = ({ advert }) => {
  const { id, title, rent, apartment } = advert;
  const [coverUrl, setCoverUrl] = useState(null);

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

  return (
    <Link to={`/advert/${id}`} className={styles.advertLink}>
      <div className={styles.advertCard}>
        {coverUrl && (
          <img
            src={coverUrl}
            alt="Фото квартиры"
            className={styles.coverPhoto}
          />
        )}
        <h3 className="text-lg font-bold">{title}</h3>
        <p>
          {apartment.city}, {apartment.street} {apartment.house}
        </p>
        <p>
          Этаж: {apartment.floor} | Комнат: {apartment.rooms}
        </p>
        <p className="font-semibold">{rent.toLocaleString()} ₽/мес</p>
      </div>
    </Link>
  );
};

export default AdvertCard;
