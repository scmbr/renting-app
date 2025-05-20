import { useState, useEffect } from "react";

const CitySelector = ({ selectedCity = "", onSelect }) => {
  const [query, setQuery] = useState(selectedCity || '');
  const [suggestions, setSuggestions] = useState([]);

    useEffect(() => {
    setQuery(selectedCity || '');
  }, [selectedCity]);

  const fetchSuggestions = async (text) => {
    if (!text) return setSuggestions([]);

    try {
      const res = await fetch(
        `https://suggest-maps.yandex.ru/v1/suggest?apikey=${import.meta.env.VITE_YANDEX_API_KEY}&text=${encodeURIComponent(text)}&lang=ru_RU&types=locality`
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
    <div style={{ position: "relative" }}>
      <input
        type="text"
        placeholder="Выберите город"
        value={query}
        onChange={handleChange}
        style={{ padding: "0.5rem", fontSize: "1rem" }}
      />
      {suggestions.length > 0 && (
        <ul
          style={{
            position: "absolute",
            top: "100%",
            left: 0,
            right: 0,
            backgroundColor: "white",
            border: "1px solid #ccc",
            zIndex: 10,
            listStyle: "none",
            padding: 0,
            margin: 0,
          }}
        >
          {suggestions.map((item, idx) => (
            <li
              key={idx}
              onClick={() => handleSelect(item.title.text)}
              style={{
                padding: "0.5rem",
                cursor: "pointer",
                borderBottom: "1px solid #eee",
              }}
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
