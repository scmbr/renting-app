import React from "react";
import AdvertCard from "@/entities/advert/components/AdvertCard";
import styles from "./AdvertList.module.css";
import { useFiltersStore } from "@/stores/useFiltersStore";

const AdvertList = ({ adverts, loading, error, total }) => {
  const filters = useFiltersStore((state) => state.filters);

  if (loading) return <p className={styles.loading}>Загрузка...</p>;
  if (error) return <p className={styles.error}>{error}</p>;

  return (
    <div className={styles.container}>
      <h2>{filters.city}</h2>
      <h5>Найдено {total} объявлений</h5>
      <div className={styles.list}>
        {adverts.map((ad) => (
          <AdvertCard key={ad.id} advert={ad} />
        ))}
      </div>
    </div>
  );
};

export default AdvertList;
