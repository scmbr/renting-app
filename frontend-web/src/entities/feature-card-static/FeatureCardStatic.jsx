import React from "react";
import styles from "./FeatureCardStatic.module.css";

const FeatureCardStatic = ({ label, icon, selected }) => (
  <div className={`${styles.featureCard} ${selected ? styles.selected : ""}`}>
    <span className={styles.featureIcon}>{icon}</span>
    <span className={styles.featureLabel}>{label}</span>
  </div>
);

export default FeatureCardStatic;
