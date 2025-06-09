import React from "react";
import AdvertCard from "@/entities/advert/components/AdvertCard";
import styles from "./AdvertList.module.css";
import { useFiltersStore } from "@/stores/useFiltersStore";

const sortOptions = [
  { label: "По дате (новые)", sort_by: "created_at", order: "desc" },
  { label: "По дате (старые)", sort_by: "created_at", order: "asc" },
  { label: "По цене (возр.)", sort_by: "rent", order: "asc" },
  { label: "По цене (убыв.)", sort_by: "rent", order: "desc" },
];

const AdvertList = ({ adverts, loading, error, total, onRemoveFavorite }) => {
  const filters = useFiltersStore((state) => state.filters);
  const updateFilter = useFiltersStore((state) => state.updateFilter);

  const handleSortChange = (e) => {
    const value = e.target.value;
    const [sort_by, order] = value.split("|");
    updateFilter("sort_by", sort_by);
    updateFilter("order", order);
  };

  if (loading) return <p className={styles.loading}>Загрузка...</p>;
  if (error) return <p className={styles.error}>{error}</p>;
  if (!adverts || adverts.length === 0)
    return <p className={styles.empty}>Объявлений не найдено</p>;

  const currentSortValue = `${filters.sort_by || "created_at"}|${
    filters.order || "desc"
  }`;

  return (
    <div className={styles.container}>
      <div className={styles.headerRow}>
        <h2 className={styles.city}>{filters.city}</h2>
        <div className={styles.sortWrapper}>
          <label htmlFor="sort-select" className={styles.sortLabel}>
            Сортировка:
          </label>
          <select
            id="sort-select"
            value={currentSortValue}
            onChange={handleSortChange}
            className={styles.sortSelect}
          >
            {sortOptions.map((option) => (
              <option
                key={`${option.sort_by}_${option.order}`}
                value={`${option.sort_by}|${option.order}`}
              >
                {option.label}
              </option>
            ))}
          </select>
        </div>
      </div>

      <h5 className={styles.found}>Найдено {total} объявлений</h5>

      <div className={styles.list}>
        {adverts.map((ad) => (
          <AdvertCard
            key={ad.id}
            advert={ad}
            onRemoveFavorite={onRemoveFavorite}
          />
        ))}
      </div>
    </div>
  );
};

export default AdvertList;
