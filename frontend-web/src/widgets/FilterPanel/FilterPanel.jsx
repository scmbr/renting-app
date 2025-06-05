import React from "react";
import styles from "./FilterPanel.module.css";
import { useFiltersStore } from "@/stores/usefiltersStore";

const FilterPanel = () => {
  const { filters, updateFilter } = useFiltersStore();

  const handleChange = (e) => {
    const { name, type, value, checked } = e.target;

    if (type === "checkbox") {
      if (checked) {
        updateFilter(name, true);
      } else {
        updateFilter(name, undefined);
      }
    } else {
      updateFilter(name, value === "" ? undefined : value);
    }
  };

  return (
    <div className={styles.filterPanel}>
      <input
        type="number"
        name="price_from"
        placeholder="Мин. цена"
        value={filters.price_from || ""}
        onChange={handleChange}
        className={styles.input}
      />
      <input
        type="number"
        name="price_to"
        placeholder="Макс. цена"
        value={filters.price_to || ""}
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

      <input
        type="number"
        name="floor_min"
        placeholder="Мин. этаж"
        value={filters.floor_min || ""}
        onChange={handleChange}
        className={styles.input}
      />
      <input
        type="number"
        name="floor_max"
        placeholder="Макс. этаж"
        value={filters.floor_max || ""}
        onChange={handleChange}
        className={styles.input}
      />

      <label className={styles.checkbox}>
        <input
          type="checkbox"
          name="elevator"
          checked={filters.elevator === true}
          onChange={handleChange}
        />
        Лифт
      </label>
      <label className={styles.checkbox}>
        <input
          type="checkbox"
          name="internet"
          checked={filters.internet === true}
          onChange={handleChange}
        />
        Интернет
      </label>
      <label className={styles.checkbox}>
        <input
          type="checkbox"
          name="washing_machine"
          checked={filters.washing_machine === true}
          onChange={handleChange}
        />
        Стиралка
      </label>
      <label className={styles.checkbox}>
        <input
          type="checkbox"
          name="conditioner"
          checked={filters.conditioner === true}
          onChange={handleChange}
        />
        Кондиционер
      </label>
    </div>
  );
};

export default FilterPanel;
