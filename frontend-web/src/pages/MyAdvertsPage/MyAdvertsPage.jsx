import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "@/shared/api/axios";
import MyAdvertCard from "@/entities/my-advert/MyAdvertCard";
import ConfirmModal from "@/shared/ui/ConfirmModal";
import styles from "./MyAdvertsPage.module.css";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";

const MyAdvertsPage = () => {
  const [adverts, setAdverts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [selectedAdvertId, setSelectedAdvertId] = useState(null);
  const [isDeleting, setIsDeleting] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchAdverts = async () => {
      try {
        const res = await api.get("/my/advert");
        setAdverts(res.data.adverts || []);
      } catch (err) {
        if (err?.response?.status === 401) {
          navigate("/login");
        } else {
          setError("Не удалось загрузить объявления");
        }
      } finally {
        setLoading(false);
      }
    };

    fetchAdverts();
  }, [navigate]);

  const handleConfirmDelete = async () => {
    if (!selectedAdvertId) return;
    setIsDeleting(true);
    try {
      await api.delete(`/my/advert/${selectedAdvertId}`);
      setAdverts((prev) => prev.filter((ad) => ad.id !== selectedAdvertId));
    } catch (err) {
      alert("Не удалось удалить объявление.");
      console.error(err);
    } finally {
      setIsDeleting(false);
      setSelectedAdvertId(null);
    }
  };

  if (loading) return <p className={styles.loading}>Загрузка объявлений...</p>;
  if (error) return <p className={styles.error}>{error}</p>;

  return (
    <>
      <SubNavbar />
      <NavPanel />
      <div className={styles.container}>
        <div className={styles.header}>
          <h1 className={styles.title}>Мои объявления</h1>
          <button
            className={styles.addButton}
            onClick={() => navigate("/my/advert/add")}
            aria-label="Добавить объявление"
          >
            <img src={"/icons/add.svg"} alt="Добавить" />
          </button>
        </div>
        {adverts.map((ad) => (
          <MyAdvertCard
            key={ad.id}
            advert={ad}
            onEdit={(id) => navigate(`/my/advert/edit/${id}`)}
            onDelete={(id) => setSelectedAdvertId(id)}
            isDeleting={isDeleting && selectedAdvertId === ad.id}
          />
        ))}
      </div>

      <ConfirmModal
        isOpen={selectedAdvertId !== null}
        onConfirm={handleConfirmDelete}
        onCancel={() => setSelectedAdvertId(null)}
        message="Вы уверены, что хотите удалить это объявление?"
      />
    </>
  );
};

export default MyAdvertsPage;
