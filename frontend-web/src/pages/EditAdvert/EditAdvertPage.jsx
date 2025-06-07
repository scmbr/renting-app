import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import api from "@/shared/api/axios";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";
import EditAdvertForm from "@/features/edit-form/ui/EditAdvertForm";
import styles from "./EditAdvertPage.module.css";
import { fetchAdvertById } from "@/entities/advert/model";
const EditAdvertPage = () => {
  const [advert, setAdvert] = useState(null);
  const navigate = useNavigate();
  const { id } = useParams();

  useEffect(() => {
    const fetchAdvert = async () => {
      try {
        const data = await fetchAdvertById(id);
        setAdvert(data);
      } catch (err) {
        console.error("Ошибка при загрузке объявления", err);
        if (err.response?.status === 401) {
          navigate("/login");
        }
      }
    };

    fetchAdvert();
  }, [id, navigate]);

  if (!advert) return <p>Загрузка...</p>;

  return (
    <div>
      <SubNavbar />
      <NavPanel />
      <div className={styles.container}>
        <h1 className={styles.title}>Редактировать объявление</h1>
        <EditAdvertForm
          apartment={advert.apartment}
          advertId={id}
          initialData={advert}
        />
      </div>
    </div>
  );
};

export default EditAdvertPage;
