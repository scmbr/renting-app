import React, { useState, useRef, useEffect } from "react";
import styles from "./FilterPanel.module.css";
import { useFiltersStore } from "@/stores/useFiltersStore";
import { FaChevronDown, FaChevronUp } from "react-icons/fa";

const FilterPanel = () => {
  const { filters, updateFilter } = useFiltersStore();

  const [priceDropdownOpen, setPriceDropdownOpen] = useState(false);
  const [roomsDropdownOpen, setRoomsDropdownOpen] = useState(false);
  const [floorDropdownOpen, setFloorDropdownOpen] = useState(false);
  const [amenitiesDropdownOpen, setAmenitiesDropdownOpen] = useState(false);

  const priceDropdownRef = useRef(null);
  const roomsDropdownRef = useRef(null);
  const floorDropdownRef = useRef(null);
  const amenitiesDropdownRef = useRef(null);

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (
        priceDropdownRef.current &&
        !priceDropdownRef.current.contains(event.target)
      )
        setPriceDropdownOpen(false);
      if (
        roomsDropdownRef.current &&
        !roomsDropdownRef.current.contains(event.target)
      )
        setRoomsDropdownOpen(false);
      if (
        floorDropdownRef.current &&
        !floorDropdownRef.current.contains(event.target)
      )
        setFloorDropdownOpen(false);
      if (
        amenitiesDropdownRef.current &&
        !amenitiesDropdownRef.current.contains(event.target)
      )
        setAmenitiesDropdownOpen(false);
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const handleChange = (e) => {
    const { name, type, value, checked } = e.target;
    if (type === "checkbox") {
      updateFilter(name, checked ? true : undefined);
    } else {
      updateFilter(name, value === "" ? undefined : value);
    }
  };

  const roomsOptions = ["1", "2", "3", "4"];
  const amenities = [
    { name: "elevator", label: "Лифт" },
    { name: "internet", label: "Интернет" },
    { name: "washing_machine", label: "Стиралка" },
    { name: "conditioner", label: "Кондиционер" },
    { name: "pets", label: "Можно с животными" },
    { name: "babies", label: "Можно с детьми" },
    { name: "smoking", label: "Можно курить" },
    { name: "tv", label: "Телевизор" },
    { name: "garbage_chute", label: "Мусоропровод" },
  ];

  return (
    <div className={styles.filterPanel}>
      <div className={styles.dropdown} ref={priceDropdownRef}>
        <button
          type="button"
          onClick={() => setPriceDropdownOpen((prev) => !prev)}
          className={styles.dropdownToggle}
        >
          Стоимость аренды{" "}
          {priceDropdownOpen ? <FaChevronUp /> : <FaChevronDown />}
        </button>
        {priceDropdownOpen && (
          <div className={styles.dropdownMenu}>
            <input
              type="number"
              name="price_from"
              placeholder="Мин. цена"
              value={filters.price_from || ""}
              onChange={handleChange}
              className={styles.input}
            />
            <input
              type="number"
              name="price_to"
              placeholder="Макс. цена"
              value={filters.price_to || ""}
              onChange={handleChange}
              className={styles.input}
            />
          </div>
        )}
      </div>

      <div className={styles.dropdown} ref={roomsDropdownRef}>
        <button
          type="button"
          onClick={() => setRoomsDropdownOpen((prev) => !prev)}
          className={styles.dropdownToggle}
        >
          {filters.rooms
            ? filters.rooms === "4"
              ? "4+ комнаты"
              : `${filters.rooms} комнат${filters.rooms === "1" ? "а" : "ы"}`
            : "Кол-во комнат"}{" "}
          {roomsDropdownOpen ? <FaChevronUp /> : <FaChevronDown />}
        </button>
        {roomsDropdownOpen && (
          <div className={styles.dropdownMenu}>
            {roomsOptions.map((room) => (
              <button
                key={room}
                type="button"
                className={styles.dropdownItem}
                onClick={() => {
                  updateFilter("rooms", room);
                  setRoomsDropdownOpen(false);
                }}
              >
                {room === "4"
                  ? "4+ комнаты"
                  : `${room} комнат${room === "1" ? "а" : "ы"}`}
              </button>
            ))}
            {filters.rooms && (
              <button
                type="button"
                className={styles.dropdownItem}
                onClick={() => {
                  updateFilter("rooms", undefined);
                  setRoomsDropdownOpen(false);
                }}
              >
                Сбросить
              </button>
            )}
          </div>
        )}
      </div>

      <div className={styles.dropdown} ref={floorDropdownRef}>
        <button
          type="button"
          onClick={() => setFloorDropdownOpen((prev) => !prev)}
          className={styles.dropdownToggle}
        >
          Этаж {floorDropdownOpen ? <FaChevronUp /> : <FaChevronDown />}
        </button>
        {floorDropdownOpen && (
          <div className={styles.dropdownMenu}>
            <input
              type="number"
              name="floor_min"
              placeholder="Мин. этаж"
              value={filters.floor_min || ""}
              onChange={handleChange}
              className={styles.input}
            />
            <input
              type="number"
              name="floor_max"
              placeholder="Макс. этаж"
              value={filters.floor_max || ""}
              onChange={handleChange}
              className={styles.input}
            />
          </div>
        )}
      </div>

      <div className={styles.dropdown} ref={amenitiesDropdownRef}>
        <button
          type="button"
          onClick={() => setAmenitiesDropdownOpen((prev) => !prev)}
          className={styles.dropdownToggle}
        >
          Удобства {amenitiesDropdownOpen ? <FaChevronUp /> : <FaChevronDown />}
        </button>
        {amenitiesDropdownOpen && (
          <div className={styles.dropdownMenu}>
            {amenities.map(({ name, label }) => (
              <label key={name} className={styles.checkbox}>
                <input
                  type="checkbox"
                  name={name}
                  checked={filters[name] === true}
                  onChange={handleChange}
                />
                {label}
              </label>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default FilterPanel;
