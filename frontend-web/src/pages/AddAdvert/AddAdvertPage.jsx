import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import AddAdvertForm from "@/features/add-form/ui/AddAdvertForm";
import api from "@/shared/api/axios";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
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

      <AddAdvertForm apartments={apartments} />
    </div>
  );
};

export default AddAdvertPage;
