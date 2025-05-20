import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import Navbar from "@/widgets/Navbar/Navbar.jsx";
import AdvertList from '@/widgets/AdvertList/AdvertList.jsx';
import { getCityFromSlug, slugify } from "@/utils/slugUtils";

const HomePage = () => {
  const { citySlug } = useParams();
  const navigate = useNavigate();
  const [selectedCity, setSelectedCity] = useState('');
  const [filters, setFilters] = useState({});


  useEffect(() => {
    if (!citySlug) {
      navigate('/moskva', { replace: true });
      return;
    }
    const cityName = getCityFromSlug(citySlug);
    setSelectedCity(cityName);
  }, [citySlug, navigate]);

  useEffect(() => {
    const newFilters = selectedCity.trim() ? { city: selectedCity.trim() } : {};
    setFilters(newFilters);
  }, [selectedCity]);

  const handleCitySelect = (city) => {
    const slug = slugify(city);
    if (slug !== citySlug) {
      navigate(`/${slug}`);
    }
    setSelectedCity(city);
  };

  return (
    <div>
      <Navbar selectedCity={selectedCity} onCitySelect={handleCitySelect} />
      <div className="p-4">
        <AdvertList filters={filters} />
      </div>
    </div>
  );
};

export default HomePage;
