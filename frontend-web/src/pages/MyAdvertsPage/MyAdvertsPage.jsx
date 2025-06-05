import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "@/shared/api/axios";
import MyAdvertCard from "@/entities/my-advert/MyAdvertCard";
import styles from "./MyAdvertsPage.module.css";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";
const MyAdvertsPage = () => {
  const [adverts, setAdverts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    let isMounted = true;

    const fetchAdverts = async () => {
      try {
        const res = await api.get("/my/advert");
        if (isMounted) {
          setAdverts(res.data || []);
          setLoading(false);
        }
      } catch (err) {
        console.error("Ошибка запроса:", err);
        if (err?.response?.status === 401) {
          navigate("/login");
        } else {
          setError("Не удалось загрузить объявления");
          setLoading(false);
        }
      }
    };

    fetchAdverts();

    return () => {
      isMounted = false;
    };
  }, [navigate]);

  if (loading) return <p className={styles.loading}>Загрузка объявлений...</p>;
  if (error) return <p className={styles.error}>{error}</p>;
  if (adverts.length === 0)
    return <p className={styles.empty}>У вас пока нет объявлений.</p>;

  return (
    <>
      <SubNavbar />
      <NavPanel />
      <div className={styles.container}>
        <h1 className={styles.title}>Мои объявления</h1>
        {adverts.map((ad) => (
          <MyAdvertCard
            key={ad.id}
            advert={ad}
            onEdit={(id) => console.log("Редактировать", id)}
            onDelete={(id) => console.log("Удалить", id)}
          />
        ))}
      </div>
    </>
  );
};

export default MyAdvertsPage;
