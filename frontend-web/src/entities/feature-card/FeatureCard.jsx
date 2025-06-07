import React from "react";
import clsx from "clsx";
import styles from "./FeatureCard.module.css";

const FeatureCard = ({ name, label, icon, selected, onClick }) => {
  return (
    <div
      className={clsx(styles.card, selected && styles.active)}
      onClick={onClick}
    >
      <span className={styles.icon}>{icon}</span>
      <span className={styles.label}>{label}</span>
    </div>
  );
};

export default FeatureCard;
