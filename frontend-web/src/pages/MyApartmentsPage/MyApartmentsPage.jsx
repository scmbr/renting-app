import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./MyApartmentsPage.module.css";
import api from "@/shared/api/axios";
import MyApartmentCard from "@/entities/my-apartment/MyApartmentCard";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";
import ConfirmModal from "@/shared/ui/ConfirmModal";

const MyApartmentPage = () => {
  const [apartments, setApartments] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [selectedApartmentId, setSelectedApartmentId] = useState(null);
  const [isDeleting, setIsDeleting] = useState(false);

  const navigate = useNavigate();

  useEffect(() => {
    const fetchApartments = async () => {
      try {
        const response = await api.get("/my/apartment");
        setApartments(response.data.apartments || []);
      } catch (err) {
        setError("Ошибка при загрузке квартир.");
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchApartments();
  }, []);

  const handleConfirmDelete = async () => {
    if (!selectedApartmentId) return;
    setIsDeleting(true);
    try {
      await api.delete(`/my/apartment/${selectedApartmentId}`);
      setApartments((prev) =>
        prev.filter((apt) => apt.id !== selectedApartmentId)
      );
    } catch (err) {
      alert("Не удалось удалить квартиру.");
      console.error(err);
    } finally {
      setIsDeleting(false);
      setSelectedApartmentId(null);
    }
  };

  if (loading) return <p className={styles.message}>Загрузка...</p>;
  if (error) return <p className={styles.message}>{error}</p>;

  return (
    <>
      <SubNavbar />
      <NavPanel />
      <div className={styles.container}>
        <div className={styles.header}>
          <h1 className={styles.title}>Мои квартиры</h1>
          <button
            className={styles.addButton}
            onClick={() => navigate("/my/apartment/add")}
            aria-label="Добавить квартиру"
          >
            <img src="/icons/add.svg" alt="Добавить" />
          </button>
        </div>
        <div className={styles.list}>
          {apartments.map((apt) => (
            <MyApartmentCard
              key={apt.id}
              apartment={apt}
              onEdit={(id) => navigate(`/my/apartment/edit/${id}`)}
              onDelete={(id) => setSelectedApartmentId(id)}
              isDeleting={isDeleting && selectedApartmentId === apt.id}
            />
          ))}
        </div>
      </div>

      <ConfirmModal
        isOpen={selectedApartmentId !== null}
        onConfirm={handleConfirmDelete}
        onCancel={() => setSelectedApartmentId(null)}
        message="Вы уверены, что хотите удалить эту квартиру?"
      />
    </>
  );
};

export default MyApartmentPage;
