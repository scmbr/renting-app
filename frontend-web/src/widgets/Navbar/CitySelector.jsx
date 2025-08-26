import { useState, useEffect } from "react";
import styles from "./CitySelector.module.css";
import { useCityStore } from "@/stores/useCityStore";
import { useFiltersStore } from "@/stores/useFiltersStore";
const CitySelector = () => {
  const city = useCityStore((state) => state.city);
  const setCity = useCityStore((state) => state.setCity);
  const [query, setQuery] = useState(city || "");
  const [suggestions, setSuggestions] = useState([]);
  const updateFilter = useFiltersStore((state) => state.updateFilter);
  const filters = useFiltersStore((state) => state.filters);
  const setFilters = useFiltersStore((state) => state.setFilters);

  useEffect(() => {
    setQuery(city || "");
  }, [city]);

  const fetchSuggestions = async (text) => {
    if (!text) return setSuggestions([]);
    try {
      const res = await fetch(
        `https://suggest-maps.yandex.ru/v1/suggest?apikey=${
          import.meta.env.VITE_YANDEX_GEOSUGGEST_API_KEY
        }&text=${encodeURIComponent(text)}&lang=ru_RU&types=locality`
      );
      const data = await res.json();
      setSuggestions(data.results || []);
    } catch (err) {
      console.error("Ошибка получения подсказок:", err);
    }
  };

  const handleChange = (e) => {
    const text = e.target.value;
    setQuery(text);
    fetchSuggestions(text);
  };

  const handleSelect = (cityName) => {
    setQuery(cityName);
    setSuggestions([]);
    setCity(cityName);
    updateFilter("city", cityName);
  };

  return (
    <div className={styles.wrapper}>
      <input
        type="text"
        placeholder="Выберите город"
        value={query}
        onChange={handleChange}
        className={styles.input}
      />
      {suggestions.length > 0 && (
        <ul className={styles.suggestions}>
          {suggestions.map((item, idx) => (
            <li
              key={idx}
              onClick={() => handleSelect(item.title.text)}
              className={styles.suggestionItem}
            >
              {item.title.text}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default CitySelector;
