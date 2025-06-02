import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "@/shared/api/axios";

const MyAdvertsPage = () => {
  const [adverts, setAdverts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    let isMounted = true;
    console.log("Загрузка объявлений...");

    const fetchAdverts = async () => {
      try {
        const res = await api.get("/my/advert");
        console.log("Ответ от API:", res.data);
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
  if (loading) return <p>Загрузка объявлений...</p>;
  if (error) return <p style={{ color: "red" }}>{error}</p>;
  if (adverts.length === 0) return <p>У вас пока нет объявлений.</p>;

  return (
    <div>
      <h1>Мои объявления</h1>
      <ul>
        {adverts.map((ad) => (
          <li key={ad.id} style={{ marginBottom: "1em" }}>
            <strong>{ad.title}</strong> — {ad.rent} ₽ / {ad.rental_type}
            <br />
            Адрес: {ad.city}, {ad.street} {ad.house}
            <br />
            Депозит: {ad.deposit} ₽
          </li>
        ))}
      </ul>
    </div>
  );
};

export default MyAdvertsPage;
