import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import api from "@/shared/api/axios";
import styles from "./AddApartmentPage.module.css";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import { MapGLForm } from "@/widgets/Map/MapGLForm";
import AddressSuggester from "@/features/add-apartment/AddressSuggester.jsx";

const AddApartmentPage = () => {
  const navigate = useNavigate();

  const [form, setForm] = useState({
    city: "",
    street: "",
    building: "",
    floor: 1,
    longitude: 0,
    latitude: 0,
    rooms: 1,
    area: "",
    elevator: false,
    garbage_chute: false,
    bathroom_type: "",
    concierge: false,
    construction_year: new Date().getFullYear(),
    construction_type: "",
    remont: "",
  });

  const [photos, setPhotos] = useState([]);
  const [previewUrls, setPreviewUrls] = useState([]);
  const [error, setError] = useState(null);
  const [cityLocation, setCityLocation] = useState(null);

  const [mapCenter, setMapCenter] = useState([37.620393, 55.75396]);
  const [mapMarker, setMapMarker] = useState([37.620393, 55.75396]);

  useEffect(() => {
    const city = localStorage.getItem("city") || "Москва";

    const fetchCoords = async () => {
      try {
        const API_KEY = import.meta.env.VITE_2GIS_MAP_API_KEY;
        const res = await fetch(
          `https://catalog.api.2gis.com/3.0/items/geocode?q=${encodeURIComponent(
            city
          )}&fields=items.point&key=${API_KEY}`
        );
        const data = await res.json();
        const point = data.result?.items?.[0]?.point;
        if (point) {
          setCityLocation(`${point.lon},${point.lat}`);
          setMapCenter([point.lon, point.lat]);
          setMapMarker([point.lon, point.lat]);
          setForm((prev) => ({
            ...prev,
            city,
            longitude: point.lon,
            latitude: point.lat,
          }));
        }
      } catch (e) {
        console.error("Ошибка геокодирования города", e);
      }
    };

    fetchCoords();
  }, []);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
    }));
  };

  const handlePhotoChange = (e) => {
    const files = Array.from(e.target.files);
    setPhotos(files);
    setPreviewUrls(files.map((file) => URL.createObjectURL(file)));
  };

  const handleAddressChange = ({
    address,
    street,
    building,
    latitude,
    longitude,
  }) => {
    setForm((prev) => ({
      ...prev,
      street,
      building,
      latitude,
      longitude,
    }));

    setMapCenter([longitude, latitude]);
    setMapMarker([longitude, latitude]);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);

    const payload = {
      city: form.city || localStorage.getItem("city") || "Москва",
      street: form.street.trim(),
      building: form.building.trim(),
      floor: Number(form.floor),
      longitude: Number(form.longitude),
      latitude: Number(form.latitude),
      rooms: Number(form.rooms),
      area: form.area ? Number(form.area) : 0,
      elevator: form.elevator,
      garbage_chute: form.garbage_chute,
      bathroom_type: form.bathroom_type.trim(),
      concierge: form.concierge,
      construction_year: Number(form.construction_year),
      construction_type: form.construction_type.trim(),
      remont: form.remont.trim(),
    };

    if (
      !payload.city ||
      !payload.street ||
      !payload.floor ||
      !payload.longitude ||
      !payload.latitude ||
      !payload.rooms ||
      !payload.construction_year
    ) {
      setError("Заполните все обязательные поля");
      return;
    }

    try {
      const res = await api.post("my/apartment", payload);
      const apartmentId = res.data.id;
      console.log("Создана квартира, ответ сервера:", res.data);
      if (photos.length > 0) {
        const photoForm = new FormData();
        photos.forEach((file) => photoForm.append("photos", file));

        await api.post(`/my/apartment/${apartmentId}/photos/batch`, photoForm, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });
      }

      navigate("/advert/add");
    } catch (err) {
      console.error(err);
      setError("Ошибка при создании квартиры или загрузке фото");
    }
  };

  return (
    <>
      <SubNavbar />
      <div className={styles.wrapper}>
        <h2 className={styles.title}>Добавление квартиры</h2>
        {error && <p className={styles.error}>{error}</p>}

        <form onSubmit={handleSubmit} className={styles.form}>
          <fieldset className={styles.section}>
            <legend>Адрес</legend>

            {cityLocation ? (
              <AddressSuggester
                className={styles.addressSuggester}
                location={cityLocation}
                onSelect={({
                  address,
                  street,
                  building,
                  latitude,
                  longitude,
                }) => {
                  setForm((prev) => ({
                    ...prev,
                    city: prev.city || localStorage.getItem("city") || "Москва",
                    street,
                    building,
                    latitude,
                    longitude,
                  }));

                  setMapCenter([longitude, latitude]);
                  setMapMarker([longitude, latitude]);
                }}
              />
            ) : (
              <p>Загрузка адресного ввода...</p>
            )}

            <div className={styles.mapContainer}>
              <MapGLForm
                center={mapCenter}
                markerPosition={mapMarker}
                onSelect={({ lng, lat }) => {
                  setForm((prev) => ({
                    ...prev,
                    latitude: lat,
                    longitude: lng,
                  }));
                  setMapCenter([lng, lat]);
                  setMapMarker([lng, lat]);
                }}
              />
            </div>
          </fieldset>

          <fieldset className={styles.section}>
            <legend>О доме</legend>
            <input
              name="floor"
              type="number"
              value={form.floor}
              onChange={handleChange}
              required
              className={styles.input}
              placeholder="Этаж"
            />
            <input
              name="construction_year"
              type="number"
              value={form.construction_year}
              onChange={handleChange}
              required
              className={styles.input}
              placeholder="Год постройки"
            />
            <input
              name="construction_type"
              value={form.construction_type}
              onChange={handleChange}
              className={styles.input}
              placeholder="Тип постройки (панель/кирпич и т.д.)"
            />
          </fieldset>

          <fieldset className={styles.section}>
            <legend>О квартире</legend>
            <input
              name="rooms"
              type="number"
              value={form.rooms}
              onChange={handleChange}
              required
              className={styles.input}
              placeholder="Количество комнат"
            />
            <input
              name="area"
              type="number"
              value={form.area}
              onChange={handleChange}
              className={styles.input}
              placeholder="Площадь (кв.м)"
            />
            <input
              name="bathroom_type"
              value={form.bathroom_type}
              onChange={handleChange}
              className={styles.input}
              placeholder="Тип санузла"
            />
            <input
              name="remont"
              value={form.remont}
              onChange={handleChange}
              className={styles.input}
              placeholder="Ремонт (евро, косметика и т.д.)"
            />
          </fieldset>

          <fieldset className={styles.section}>
            <legend>Дополнительно</legend>
            <label className={styles.checkbox}>
              <input
                type="checkbox"
                name="elevator"
                checked={form.elevator}
                onChange={handleChange}
              />{" "}
              Лифт
            </label>
            <label className={styles.checkbox}>
              <input
                type="checkbox"
                name="garbage_chute"
                checked={form.garbage_chute}
                onChange={handleChange}
              />{" "}
              Мусоропровод
            </label>
            <label className={styles.checkbox}>
              <input
                type="checkbox"
                name="concierge"
                checked={form.concierge}
                onChange={handleChange}
              />{" "}
              Консьерж
            </label>
          </fieldset>

          <fieldset className={styles.section}>
            <legend>Фотографии</legend>
            <label className={styles.fileInput}>
              <span>Загрузите фото:</span>
              <input
                type="file"
                multiple
                accept="image/*"
                onChange={handlePhotoChange}
              />
            </label>
            <div className={styles.previewContainer}>
              {previewUrls.map((url, idx) => (
                <img
                  key={idx}
                  src={url}
                  alt={`preview-${idx}`}
                  className={styles.previewImage}
                />
              ))}
            </div>
          </fieldset>

          <button type="submit" className={styles.button}>
            Добавить квартиру
          </button>
        </form>
      </div>
    </>
  );
};

export default AddApartmentPage;
