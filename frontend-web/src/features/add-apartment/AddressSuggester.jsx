import { useState, useEffect, useRef } from "react";
import styles from "./AddressSuggester.module.css";

const GEOSUGGEST_API_KEY = import.meta.env.VITE_YANDEX_GEOSUGGEST_API_KEY;
const GEOCODER_API_KEY = import.meta.env.VITE_YANDEX_GEOCODER_KEY;

const AddressSuggester = ({
  onSelect,
  placeholder = "Введите адрес",
  value = "",
}) => {
  const [query, setQuery] = useState(value ?? "");
  const [suggestions, setSuggestions] = useState([]);
  const wrapperRef = useRef(null);

  useEffect(() => {
    setQuery(value ?? "");
  }, [value]);

  useEffect(() => {
    const handleClickOutside = (e) => {
      if (wrapperRef.current && !wrapperRef.current.contains(e.target)) {
        setSuggestions([]);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const fetchSuggestions = async (text) => {
    const cityFromStorage = localStorage.getItem("city") || "Москва";
    if (!text) return setSuggestions([]);
    try {
      const query = `${cityFromStorage} ${text}`;
      const url = `https://suggest-maps.yandex.ru/v1/suggest?apikey=${GEOSUGGEST_API_KEY}&text=${encodeURIComponent(
        query
      )}&lang=ru_RU&types=house`;
      console.log("Suggest URL:", url);

      const res = await fetch(url);
      const data = await res.json();

      console.log("Ответ SUGGEST:", data);
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

  const handleSelect = async (item) => {
    const address = `${item.title?.text || ""}`.trim();
    setQuery(address);
    setSuggestions([]);

    try {
      const res = await fetch(
        `https://geocode-maps.yandex.ru/1.x/?apikey=${GEOCODER_API_KEY}&format=json&geocode=${encodeURIComponent(
          address
        )}`
      );
      const data = await res.json();
      console.log("Ответ геокодера:", data);

      const pos =
        data.response?.GeoObjectCollection?.featureMember?.[0]?.GeoObject?.Point
          ?.pos || "0 0";
      const [longitude, latitude] = pos.split(" ").map(Number);

      onSelect({
        address,
        street: address,
        building: "",
        longitude,
        latitude,
      });
    } catch (err) {
      console.error("Ошибка получения координат:", err);
      onSelect({
        address,
        street: address,
        building: "",
        longitude: 0,
        latitude: 0,
      });
    }
  };

  return (
    <div className={styles.wrapper} ref={wrapperRef}>
      <input
        type="text"
        placeholder={placeholder}
        value={query || ""}
        onChange={handleChange}
        className={styles.input}
      />
      {suggestions.length > 0 && (
        <ul className={styles.suggestions}>
          {suggestions.map((item, idx) => (
            <li
              key={idx}
              onClick={() => handleSelect(item)}
              className={styles.suggestionItem}
            >
              {item.title?.text || "Без названия"}
              {item.subtitle?.text ? `, ${item.subtitle.text}` : ""}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default AddressSuggester;
