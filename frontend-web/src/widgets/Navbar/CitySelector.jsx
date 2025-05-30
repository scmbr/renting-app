import { useState, useEffect } from "react";
import styles from "./CitySelector.module.css";

const CitySelector = ({ selectedCity = "", onSelect }) => {
  const [query, setQuery] = useState(selectedCity || "");
  const [suggestions, setSuggestions] = useState([]);

  useEffect(() => {
    setQuery(selectedCity || "");
  }, [selectedCity]);

  const fetchSuggestions = async (text) => {
    if (!text) return setSuggestions([]);
    try {
      const res = await fetch(
        `https://suggest-maps.yandex.ru/v1/suggest?apikey=${
          import.meta.env.VITE_YANDEX_API_KEY
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

  const handleSelect = (city) => {
    setQuery(city);
    setSuggestions([]);
    onSelect(city);
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
