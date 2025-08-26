import { useState, useEffect, useRef, useCallback } from "react";
import styles from "./AddressSuggester.module.css";

const GEOSUGGEST_API_KEY = import.meta.env.VITE_YANDEX_GEOSUGGEST_API_KEY;
const GEOCODER_API_KEY = import.meta.env.VITE_YANDEX_GEOCODER_KEY;

const sanitizeText = (text) => text?.replace(/<[^>]*>?/gm, "") || "";

const AddressSuggester = ({
  onSelect,
  placeholder = "Введите адрес",
  value = "",
}) => {
  const [query, setQuery] = useState(value ?? "");
  const [suggestions, setSuggestions] = useState([]);
  const [highlightIndex, setHighlightIndex] = useState(-1);
  const [loading, setLoading] = useState(false);

  const wrapperRef = useRef(null);
  const debounceRef = useRef(null);
  const activeRef = useRef(true);

  useEffect(() => {
    setQuery(value ?? "");
  }, [value]);

  useEffect(() => {
    activeRef.current = true;
    return () => {
      activeRef.current = false;
    };
  }, []);

  useEffect(() => {
    const handleClickOutside = (e) => {
      if (wrapperRef.current && !wrapperRef.current.contains(e.target)) {
        setSuggestions([]);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const fetchSuggestions = useCallback(async (text) => {
    const cityFromStorage = localStorage.getItem("city") || "Москва";
    if (!text.trim()) return setSuggestions([]);

    setLoading(true);
    try {
      const url = `https://suggest-maps.yandex.ru/v1/suggest?apikey=${GEOSUGGEST_API_KEY}&text=${encodeURIComponent(
        `${cityFromStorage} ${text}`
      )}&lang=ru_RU&types=house`;

      const res = await fetch(url);
      if (!res.ok) throw new Error(`Ошибка: ${res.status}`);
      const data = await res.json();

      if (activeRef.current) {
        setSuggestions(data.results || []);
      }
    } catch (err) {
      console.error("Ошибка получения подсказок:", err);
      if (activeRef.current) setSuggestions([]);
    } finally {
      if (activeRef.current) setLoading(false);
    }
  }, []);

  const handleChange = (e) => {
    const text = e.target.value;
    setQuery(text);
    clearTimeout(debounceRef.current);
    debounceRef.current = setTimeout(() => fetchSuggestions(text), 400);
  };

  const handleSelect = async (item) => {
    const address = sanitizeText(item.title?.text || "");
    setQuery(address);
    setSuggestions([]);
    setHighlightIndex(-1);

    try {
      const res = await fetch(
        `https://geocode-maps.yandex.ru/1.x/?apikey=${GEOCODER_API_KEY}&format=json&geocode=${encodeURIComponent(
          address
        )}`
      );
      if (!res.ok) throw new Error(`Ошибка геокодера: ${res.status}`);

      const data = await res.json();
      const pos =
        data.response?.GeoObjectCollection?.featureMember?.[0]?.GeoObject?.Point
          ?.pos || "0 0";
      const [longitude, latitude] = pos.split(" ").map(Number);

      onSelect({ address, street: address, building: "", longitude, latitude });
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

  const handleKeyDown = (e) => {
    if (!suggestions.length) return;

    if (e.key === "ArrowDown") {
      e.preventDefault();
      setHighlightIndex((prev) =>
        prev < suggestions.length - 1 ? prev + 1 : 0
      );
    } else if (e.key === "ArrowUp") {
      e.preventDefault();
      setHighlightIndex((prev) =>
        prev > 0 ? prev - 1 : suggestions.length - 1
      );
    } else if (e.key === "Enter" && highlightIndex >= 0) {
      e.preventDefault();
      handleSelect(suggestions[highlightIndex]);
    }
  };

  return (
    <div className={styles.wrapper} ref={wrapperRef}>
      <input
        type="text"
        placeholder={placeholder}
        value={query}
        onChange={handleChange}
        onKeyDown={handleKeyDown}
        className={styles.input}
      />
      {loading && <div className={styles.loader}>Загрузка...</div>}
      {!loading && suggestions.length > 0 && (
        <ul className={styles.suggestions}>
          {suggestions.map((item, idx) => (
            <li
              key={idx}
              onClick={() => handleSelect(item)}
              className={`${styles.suggestionItem} ${
                idx === highlightIndex ? styles.highlighted : ""
              }`}
            >
              {sanitizeText(item.title?.text) || "Без названия"}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default AddressSuggester;
