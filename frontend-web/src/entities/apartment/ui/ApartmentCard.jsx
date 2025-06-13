import React from "react";
import styles from "./ApartmentCard.module.css";

const ApartmentCard = ({ apartment, isSelected, onSelect }) => {
  const hasImage = apartment.images && apartment.images.length > 0;

  return (
    <div
      className={`${styles.apartmentCard} ${isSelected ? styles.selected : ""}`}
      onClick={() => onSelect(apartment.id)}
    >
      {hasImage && (
        <img
          src={apartment.images[0]}
          alt="apartment"
          className={styles.apartmentImage}
        />
      )}
      <div className={styles.apartmentInfo}>
        <p>
          {apartment.city}, {apartment.street} {apartment.building}
        </p>
      </div>
    </div>
  );
};

export default ApartmentCard;
