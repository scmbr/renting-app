import { useState, useEffect, useRef } from "react";
import styles from "./AddressSuggester.module.css";

const API_KEY = import.meta.env.VITE_2GIS_MAP_API_KEY;

const AddressSuggester = ({
  onSelect,
  placeholder = "Введите адрес",
  location,
  value = "",
}) => {
  const [query, setQuery] = useState("");
  const [suggestions, setSuggestions] = useState([]);
  const wrapperRef = useRef(null);

  useEffect(() => {
    setQuery(value || "");
  }, [value]);
  useEffect(() => {
    if (!location) setSuggestions([]);
  }, [location]);

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
    if (!text || !location) return setSuggestions([]);
    try {
      const url = `https://catalog.api.2gis.com/3.0/suggests?q=${encodeURIComponent(
        text
      )}&location=${location}&key=${API_KEY}`;

      const res = await fetch(url);
      const data = await res.json();
      setSuggestions(data.result?.items || []);
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
    setQuery(item.name);
    setSuggestions([]);
    const cityFromStorage = localStorage.getItem("city") || "Москва";
    try {
      const query = `${cityFromStorage} ${item.name}`;
      const url = `https://catalog.api.2gis.com/3.0/items/geocode?q=${encodeURIComponent(
        query
      )}&fields=items.point&key=${API_KEY}`;
      const res = await fetch(url);
      const data = await res.json();

      const point = data.result?.items?.[0]?.point;
      const longitude = point?.lon || 0;
      const latitude = point?.lat || 0;

      let fullAddress = item.name.trim();

      const prefixes = [
        "улица ",
        "проспект ",
        "бульвар ",
        "переулок ",
        "шоссе ",
        "аллея ",
        "площадь ",
      ];

      for (const prefix of prefixes) {
        if (fullAddress.toLowerCase().startsWith(prefix)) {
          fullAddress = fullAddress.substring(prefix.length).trim();
          break;
        }
      }

      const parts = fullAddress.split(",").map((part) => part.trim());

      const street = parts[0] || "";
      const building = parts[1] || "";

      onSelect({
        address: item.name,
        street,
        building,
        longitude,
        latitude,
      });
    } catch (error) {
      console.error("Ошибка при получении координат:", error);

      let fullAddress = item.name.trim();

      const prefixes = [
        "улица ",
        "проспект ",
        "бульвар ",
        "переулок ",
        "шоссе ",
        "аллея ",
        "площадь ",
      ];

      for (const prefix of prefixes) {
        if (fullAddress.toLowerCase().startsWith(prefix)) {
          fullAddress = fullAddress.substring(prefix.length).trim();
          break;
        }
      }

      const parts = fullAddress.split(",").map((part) => part.trim());
      const street = parts[0] || "";
      const building = parts[1] || "";

      onSelect({
        address: item.name,
        street,
        building,
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
        value={query}
        onChange={handleChange}
        className={styles.input}
        disabled={!location}
      />
      {suggestions.length > 0 && (
        <ul className={styles.suggestions}>
          {suggestions.map((item, idx) => (
            <li
              key={idx}
              onClick={() => handleSelect(item)}
              className={styles.suggestionItem}
            >
              {item.name}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default AddressSuggester;
