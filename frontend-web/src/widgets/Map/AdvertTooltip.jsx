import React from "react";
import styles from "./AdvertTooltip.module.css";

const AdvertTooltip = ({ advert }) => {
  if (!advert || !advert.apartment) return null;

  const { title, rent, apartment, coverUrl } = advert;

  return (
    <div className={styles.tooltipCard}>
      <img
        src={coverUrl || "/images/no-photo.png"}
        alt="Фото квартиры"
        className={styles.coverPhoto}
      />
      <div className={styles.content}>
        <h4 className={styles.title}>{title}</h4>
        <p className={styles.price}>{rent.toLocaleString()} ₽/мес</p>
        <p className={styles.address}>
          {apartment.city}, {apartment.street} {apartment.house}
        </p>
      </div>
    </div>
  );
};

export default AdvertTooltip;
