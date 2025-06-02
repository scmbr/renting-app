import React from "react";
import styles from "./FilterPanel.module.css";

const FilterPanel = ({ filters, setFilters }) => {
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFilters((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  return (
    <div className={styles.filterPanel}>
      <input
        type="number"
        name="priceMin"
        placeholder="Мин. цена"
        value={filters.priceMin || ""}
        onChange={handleChange}
        className={styles.input}
      />
      <input
        type="number"
        name="priceMax"
        placeholder="Макс. цена"
        value={filters.priceMax || ""}
        onChange={handleChange}
        className={styles.input}
      />
      <select
        name="rooms"
        value={filters.rooms || ""}
        onChange={handleChange}
        className={styles.select}
      >
        <option value="">Кол-во комнат</option>
        <option value="1">1 комната</option>
        <option value="2">2 комнаты</option>
        <option value="3">3 комнаты</option>
        <option value="4">4+</option>
      </select>
    </div>
  );
};

export default FilterPanel;
