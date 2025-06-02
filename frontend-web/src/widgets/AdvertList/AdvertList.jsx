import React, { useEffect, useState } from "react";
import { fetchAdverts } from "@/entities/advert/model";
import AdvertCard from "@/entities/advert/components/AdvertCard";
import styles from "./AdvertList.module.css";

const AdvertList = ({ filters }) => {
  const [adverts, setAdverts] = useState([]);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    setLoading(true);
    setError(null);

    fetchAdverts(filters)
      .then((data) => {
        setAdverts(Array.isArray(data.adverts) ? data.adverts : []);
        setTotal(data.total ?? 0);
      })
      .catch((err) => {
        setError("Ошибка при загрузке объявлений");
        console.error(err);
      })
      .finally(() => setLoading(false));
  }, [filters]);

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
