import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import AddAdvertForm from "@/features/add-form/ui/AddAdvertForm";
import api from "@/shared/api/axios";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import styles from "./AddAdvertPage.module.css";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";
const AddAdvertPage = () => {
  const [apartments, setApartments] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    let isMounted = true;

    const fetchApartments = async () => {
      try {
        const res = await api.get("/my/apartment");
        const data = res.data;
        const apartmentList = data.apartments;

        if (isMounted) {
          setApartments(apartmentList);
        }
      } catch (err) {
        if (err.response?.status === 401) {
          navigate("/login");
        }
      }
    };

    fetchApartments();

    return () => {
      isMounted = false;
    };
  }, [navigate]);

  return (
    <div>
      <SubNavbar />
      <NavPanel />
      <div className={styles.container}>
        <h1 className={styles.title}>Добавить объявление</h1>
        <AddAdvertForm apartments={apartments} />
      </div>
    </div>
  );
};

export default AddAdvertPage;
