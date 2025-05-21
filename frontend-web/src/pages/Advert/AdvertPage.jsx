import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { fetchAdvertById } from "@/entities/advert/model";

const AdvertPage = () => {
  const { id } = useParams();
  const [advert, setAdvert] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const loadAdvert = async () => {
      console.log("Загружаем объявление с ID:", id); // <-- добавь

      try {
        const data = await fetchAdvertById(id);
        console.log("Получено объявление:", data); // <-- добавь
        setAdvert(data);
      } catch (err) {
        console.error("Ошибка при загрузке:", err); // <-- добавь
        setError("Ошибка при загрузке объявления");
      } finally {
        setLoading(false);
      }
    };

    loadAdvert();
  }, [id]);

  if (loading) return <p>Загрузка объявления...</p>;
  if (error) return <p style={{ color: "red" }}>{error}</p>;
  if (!advert) return null;

  return (
    <div style={{ maxWidth: "600px", margin: "auto", padding: "1rem" }}>
      <h1>{advert.title}</h1>
      <p>
        {advert.apartment.city}, {advert.apartment.street}{" "}
        {advert.apartment.house}
      </p>
      <p>
        Этаж: {advert.apartment.floor} | Комнат: {advert.apartment.rooms}
      </p>
      <p style={{ fontWeight: "bold", fontSize: "1.2rem" }}>
        {advert.rent.toLocaleString()} ₽/мес
      </p>
    </div>
  );
};

export default AdvertPage;
