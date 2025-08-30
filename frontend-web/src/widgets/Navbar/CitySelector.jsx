import { useState, useEffect, useRef } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import styles from "./CitySelector.module.css";
import { useCityStore } from "@/stores/useCityStore";
import { useFiltersStore } from "@/stores/useFiltersStore";

const CitySelector = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const city = useCityStore((state) => state.city);
  const setCity = useCityStore((state) => state.setCity);
  const updateFilter = useFiltersStore((state) => state.updateFilter);

  const [query, setQuery] = useState(city || ""); // показываем город, если он есть
  const [fetchQuery, setFetchQuery] = useState(""); // сюда кладем текст для fetch
  const [suggestions, setSuggestions] = useState([]);
  const [loading, setLoading] = useState(false);
  const wrapperRef = useRef();
  const abortControllerRef = useRef(null);

  // Закрытие подсказок при клике вне компонента
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (wrapperRef.current && !wrapperRef.current.contains(event.target)) {
        setSuggestions([]);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  // Fetch подсказок с дебаунсом и защитой от race condition
  useEffect(() => {
    if (!fetchQuery) {
      setSuggestions([]);
      return;
    }

    setLoading(true);

    if (abortControllerRef.current) {
      abortControllerRef.current.abort();
    }
    const controller = new AbortController();
    abortControllerRef.current = controller;

    const timeoutId = setTimeout(async () => {
      try {
        const res = await fetch(
          `https://suggest-maps.yandex.ru/v1/suggest?apikey=${
            import.meta.env.VITE_YANDEX_GEOSUGGEST_API_KEY
          }&text=${encodeURIComponent(fetchQuery)}&lang=ru_RU&types=locality`,
          { signal: controller.signal }
        );
        const data = await res.json();
        const results = data.results || [];

        if (!controller.signal.aborted) {
          setSuggestions(results);
        }
      } catch (err) {
        if (err.name !== "AbortError") {
          console.error("Ошибка получения подсказок:", err);
        }
      } finally {
        if (!controller.signal.aborted) {
          setLoading(false);
        }
      }
    }, 300);

    return () => {
      clearTimeout(timeoutId);
      controller.abort();
    };
  }, [fetchQuery]);

  const handleChange = (e) => {
    const value = e.target.value;
    setQuery(value); // для отображения в input
    setFetchQuery(value); // для fetch подсказок
  };

  const handleSelect = (cityName) => {
    if (!cityName) return;
    setQuery(cityName);
    setSuggestions([]);
    setCity(cityName);
    updateFilter("city", cityName);

    // Редирект на главную только после выбора города
    if (location.pathname !== "/") {
      navigate("/", { replace: true });
    }
  };

  return (
    <div className={styles.wrapper} ref={wrapperRef}>
      <input
        type="text"
        placeholder="Выберите город"
        value={query}
        onChange={handleChange}
        className={styles.input}
      />
      {suggestions.length > 0 && (
        <ul className={styles.suggestions}>
          {suggestions.map((item, idx) => {
            const name = item.displayName || item.title?.text || "";
            return (
              <li
                key={item.id || name || idx}
                onClick={() => handleSelect(name)}
                className={styles.suggestionItem}
              >
                {name}
              </li>
            );
          })}
        </ul>
      )}
    </div>
  );
};

export default CitySelector;
