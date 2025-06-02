import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import Navbar from "@/widgets/Navbar/Navbar.jsx";
import AdvertList from "@/widgets/AdvertList/AdvertList.jsx";
import FilterPanel from "@/widgets/FilterPanel/FilterPanel.jsx";
import { MapGL } from "@/widgets/Map/2GIS.jsx";
import { slugToName, nameToSlug } from "@/shared/constants/cities";
import styles from "./HomePage.module.css";

const HomePage = () => {
  const { citySlug } = useParams();
  const navigate = useNavigate();
  const [selectedCity, setSelectedCity] = useState("Москва");
  const [filters, setFilters] = useState({});

  useEffect(() => {
    const newFilters = selectedCity.trim() ? { city: selectedCity.trim() } : {};
    setFilters(newFilters);
  }, [selectedCity]);

  useEffect(() => {
    if (!citySlug) {
      navigate("/moskva", { replace: true });
      return;
    }
    setSelectedCity(slugToName(citySlug));
  }, [citySlug]);

  const handleCitySelect = (city) => {
    const citySlugNew = nameToSlug(city);
    navigate(`/${citySlugNew}`);
  };

  return (
    <div>
      <Navbar selectedCity={selectedCity} onCitySelect={handleCitySelect} />
      <FilterPanel filters={filters} setFilters={setFilters} />
      <div className={styles.container}>
        <div className={styles.mapContainer}>
          <MapGL />
        </div>

        <div className={styles.advertsContainer}>
          <AdvertList filters={filters} />
        </div>
      </div>
    </div>
  );
};

export default HomePage;
