import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import AdvertList from "@/widgets/AdvertList/AdvertList.jsx";
import FilterPanel from "@/widgets/FilterPanel/FilterPanel.jsx";
import { YandexMap } from "@/widgets/Map/YandexMap.jsx";
import { slugToName, nameToSlug } from "@/shared/constants/cities";
import styles from "./HomePage.module.css";
import { useEffect, useState, useCallback } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useCityStore } from "@/stores/useCityStore";
import { useFiltersStore } from "@/stores/useFiltersStore";
import { fetchAdverts } from "@/entities/advert/model";

const HomePage = () => {
  const { citySlug } = useParams();
  const navigate = useNavigate();

  const city = useCityStore((state) => state.city);
  const setCity = useCityStore((state) => state.setCity);

  const filters = useFiltersStore((state) => state.filters);
  const updateFilter = useFiltersStore((state) => state.updateFilter);

  const [adverts, setAdverts] = useState([]);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (!citySlug) {
      const storedCity = localStorage.getItem("city") || "Москва";
      setCity(storedCity);
      updateFilter("city", storedCity);
      navigate(`/${nameToSlug(storedCity)}`, { replace: true });
      return;
    }

    const cityName = slugToName(citySlug);
    setCity(cityName);
    updateFilter("city", cityName);
  }, [citySlug, navigate, setCity, updateFilter]);

  const loadAdverts = useCallback(() => {
    if (!filters.city) return;
    setLoading(true);
    setError(null);

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

  useEffect(() => {
    loadAdverts();
  }, [loadAdverts]);

  const handleCitySelect = (newCity) => {
    setCity(newCity);
    updateFilter("city", newCity);
    navigate(`/${nameToSlug(newCity)}`);
  };

  return (
    <>
      <SubNavbar selectedCity={city} onCitySelect={handleCitySelect} />
      <FilterPanel />
      <div className={styles.container}>
        <div className={styles.mapContainer}>
          {city ? (
            <YandexMap adverts={adverts} city={city} />
          ) : (
            <div>Загрузка карты...</div>
          )}
        </div>

        <div className={styles.advertsContainer}>
          <AdvertList
            adverts={adverts}
            loading={loading}
            error={error}
            total={total}
            onRemoveFavorite={() => {}}
            onRetry={loadAdverts}
          />
        </div>
      </div>
    </>
  );
};

export default HomePage;
