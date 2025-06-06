import React from "react";
import AdvertCard from "@/entities/advert/components/AdvertCard";
import styles from "./FavoritesList.module.css";
import { useFiltersStore } from "@/stores/useFiltersStore";

const AdvertList = ({ adverts, loading, error, total, onRemoveFavorite }) => {
  const filters = useFiltersStore((state) => state.filters);

  if (loading) return <p className={styles.loading}>Загрузка...</p>;
  if (error) return <p className={styles.error}>{error}</p>;

  if (!adverts || adverts.length === 0) {
    return <p className={styles.empty}>Объявлений не найдено</p>;
  }

  return (
    <div className={styles.container}>
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
