import Navbar from "@/widgets/Navbar/Navbar.jsx";
import AdvertList from "@/widgets/AdvertList/AdvertList.jsx";
import FilterPanel from "@/widgets/FilterPanel/FilterPanel.jsx";
import { MapGL } from "@/widgets/Map/2GIS.jsx";
import { slugToName, nameToSlug } from "@/shared/constants/cities";
import styles from "./HomePage.module.css";
import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useCityStore } from "@/stores/useCityStore";
import { useFiltersStore } from "@/stores/useFiltersStore";
import { fetchAdverts } from "@/entities/advert/model";

const HomePage = () => {
  const { citySlug } = useParams();
  const navigate = useNavigate();
  const city = useCityStore((state) => state.city);
  const setCity = useCityStore((state) => state.setCity);
  const updateFilter = useFiltersStore((state) => state.updateFilter);
  const filters = useFiltersStore((state) => state.filters);
  const setFilters = useFiltersStore((state) => state.setFilters);

  const [adverts, setAdverts] = useState([]);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (!citySlug) {
      const storedCity = localStorage.getItem("city");
      const fallbackCity = storedCity || "Москва";
      const slug = nameToSlug(fallbackCity);

      setCity(fallbackCity);
      updateFilter("city", fallbackCity);
      navigate(`/${slug}`, { replace: true });
    } else {
      const cityName = slugToName(citySlug);
      setCity(cityName);
      updateFilter("city", cityName);
    }
  }, [citySlug, navigate]);

  useEffect(() => {
    if (!filters.city) return;
    setLoading(true);
    fetchAdverts(filters)
      .then((data) => {
        setAdverts(Array.isArray(data.adverts) ? data.adverts : []);
        setTotal(data.total ?? 0);
      })
      .catch((err) => {
        console.error(err);
        setError("Ошибка при загрузке объявлений");
      })
      .finally(() => setLoading(false));
  }, [filters]);

  const handleCitySelect = (newCity) => {
    setCity(newCity);
    updateFilter("city", newCity);
    const newSlug = nameToSlug(newCity);
    navigate(`/${newSlug}`);
  };

  return (
    <>
      <Navbar selectedCity={city} onCitySelect={handleCitySelect} />
      <FilterPanel />
      <div className={styles.container}>
        <div className={styles.mapContainer}>
          {city ? <MapGL adverts={adverts} /> : <div>Загрузка карты...</div>}
        </div>

        <div className={styles.advertsContainer}>
          <AdvertList
            adverts={adverts}
            loading={loading}
            error={error}
            total={total}
          />
        </div>
      </div>
    </>
  );
};

export default HomePage;
