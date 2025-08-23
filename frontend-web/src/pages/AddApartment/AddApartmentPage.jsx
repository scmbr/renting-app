import React, { useState, useEffect } from "react";
import api from "@/shared/api/axios";
import styles from "./AddApartmentPage.module.css";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import { MapGLForm } from "@/widgets/Map/MapGLForm";
import AddressSuggester from "@/features/add-apartment/AddressSuggester.jsx";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";
import PhotoUploader from "@/features/upload-photo/PhotoUploader";
import { useAdvertStore } from "@/stores/useAdvertStore";
import { useNavigate, useLocation } from "react-router-dom";
const AddApartmentPage = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const from = location.state?.from || "apartment";
  const [form, setForm] = useState({
    city: "",
    street: "",
    building: "",
    floor: "",
    longitude: 0,
    latitude: 0,
    rooms: "",
    area: "",
    elevator: false,
    garbage_chute: false,
    bathroom_type: "",
    concierge: false,
    construction_year: "",
    construction_type: "",
    remont: "",
  });

  const [photos, setPhotos] = useState([]);
  const [previewUrls, setPreviewUrls] = useState([]);
  const [error, setError] = useState(null);
  const [cityLocation, setCityLocation] = useState(null);
  const { clearFilledFormData } = useAdvertStore();
  const [mapCenter, setMapCenter] = useState([37.620393, 55.75396]);
  const [mapMarker, setMapMarker] = useState([37.620393, 55.75396]);

  useEffect(() => {
    const city = localStorage.getItem("city") || "Москва";

    const fetchCoords = async () => {
      try {
        const API_KEY = import.meta.env.VITE_YANDEX_GEOCODER_KEY;
        const res = await fetch(
          `https://geocode-maps.yandex.ru/v1/?apikey=${API_KEY}&geocode=${encodeURIComponent(
            city
          )}&format=json`
        );
        const data = await res.json();
        const pos =
          data.response?.GeoObjectCollection?.featureMember?.[0]?.GeoObject
            ?.Point?.pos || "0 0";
        const [lon, lat] = pos.split(" ").map(Number);

        if (lon && lat) {
          setCityLocation(`${lat},${lon}`); // для отображения в UI
          setMapCenter([lat, lon]); // для карты
          setMapMarker([lat, lon]);
          setForm((prev) => ({
            ...prev,
            city,
            longitude: lon,
            latitude: lat,
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

    setMapCenter([latitude, longitude]);
    setMapMarker([latitude, longitude]);
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

      if (from === "advert") {
        navigate("/my/advert/add");
      } else {
        clearFilledFormData();
        navigate("/my/apartment");
      }
    } catch (err) {
      console.error(err);
      setError("Ошибка при создании квартиры или загрузке фото");
    }
  };

  return (
    <>
      <SubNavbar />
      <NavPanel />
      <div className={styles.wrapper}>
        <h2 className={styles.title}>Добавление квартиры</h2>
        {error && <p className={styles.error}>{error}</p>}

        <form onSubmit={handleSubmit} className={styles.form}>
          <div className={styles.section}>
            <h3>Адрес</h3>
            {cityLocation ? (
              <AddressSuggester
                className={styles.addressSuggester}
                location={cityLocation}
                onSelect={handleAddressChange}
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
          </div>

          <div className={styles.section}>
            <h3>О доме</h3>
            <div className={styles.inputRow}>
              <label htmlFor="floor">Этаж:</label>
              <input
                id="floor"
                name="floor"
                type="number"
                value={form.floor}
                onChange={handleChange}
                required
                className={styles.inputField}
              />
            </div>

            <div className={styles.inputRow}>
              <label htmlFor="construction_year">Год постройки:</label>
              <input
                id="construction_year"
                name="construction_year"
                type="number"
                value={form.construction_year}
                onChange={handleChange}
                required
                className={styles.inputField}
              />
            </div>

            <div className={styles.inputRowRadio}>
              <label>Тип постройки:</label>
              <div>
                <div className={styles.radioGroup}>
                  <label className={styles.customRadio}>
                    <input
                      type="radio"
                      name="construction_type"
                      value="панель"
                      checked={form.construction_type === "панель"}
                      onChange={handleChange}
                    />
                    <span className={styles.radioText}>панель</span>
                  </label>

                  <label className={styles.customRadio}>
                    <input
                      type="radio"
                      name="construction_type"
                      value="хрущевка"
                      checked={form.construction_type === "хрущевка"}
                      onChange={handleChange}
                    />
                    <span className={styles.radioText}>хрущевка</span>
                  </label>

                  <label className={styles.customRadio}>
                    <input
                      type="radio"
                      name="construction_type"
                      value="монолит"
                      checked={form.construction_type === "монолит"}
                      onChange={handleChange}
                    />
                    <span className={styles.radioText}>монолит</span>
                  </label>
                  <label className={styles.customRadio}>
                    <input
                      type="radio"
                      name="construction_type"
                      value="кирпич"
                      checked={form.construction_type === "кирпич"}
                      onChange={handleChange}
                    />
                    <span className={styles.radioText}>кирпич</span>
                  </label>
                </div>
              </div>
            </div>
          </div>

          <div className={styles.section}>
            <h3>О квартире</h3>
            <div className={styles.inputRow}>
              <label htmlFor="rooms">Количество комнат:</label>
              <input
                id="rooms"
                name="rooms"
                type="number"
                value={form.rooms}
                onChange={handleChange}
                required
                className={styles.inputField}
              />
            </div>
            <div className={styles.inputRow}>
              <label htmlFor="area">Площадь (кв.м):</label>
              <input
                id="area"
                name="area"
                type="number"
                value={form.area}
                onChange={handleChange}
                className={styles.inputField}
              />
            </div>
            <div className={styles.inputRowRadio}>
              <label>Тип санузла:</label>
              <div>
                <div className={styles.radioGroup}>
                  <label className={styles.customRadio}>
                    <input
                      type="radio"
                      name="bathroom_type"
                      value="раздельный"
                      checked={form.bathroom_type === "раздельный"}
                      onChange={handleChange}
                    />
                    <span className={styles.radioText}>раздельный</span>
                  </label>

                  <label className={styles.customRadio}>
                    <input
                      type="radio"
                      name="bathroom_type"
                      value="совмещенный"
                      checked={form.bathroom_type === "совмещенный"}
                      onChange={handleChange}
                    />
                    <span className={styles.radioText}>совмещенный</span>
                  </label>

                  <label className={styles.customRadio}>
                    <input
                      type="radio"
                      name="bathroom_type"
                      value="более одного"
                      checked={form.bathroom_type === "более одного"}
                      onChange={handleChange}
                    />
                    <span className={styles.radioText}>более одного</span>
                  </label>
                </div>
              </div>
            </div>

            <div className={styles.inputRowRadio}>
              <label>Тип ремонта:</label>
              <div>
                <div className={styles.radioGroup}>
                  <label className={styles.customRadio}>
                    <input
                      type="radio"
                      name="remont_type"
                      value="косметический"
                      checked={form.remont_type === "косметический"}
                      onChange={handleChange}
                    />
                    <span className={styles.radioText}>косметический</span>
                  </label>

                  <label className={styles.customRadio}>
                    <input
                      type="radio"
                      name="remont_type"
                      value="евро"
                      checked={form.remont_type === "евро"}
                      onChange={handleChange}
                    />
                    <span className={styles.radioText}>евро</span>
                  </label>

                  <label className={styles.customRadio}>
                    <input
                      type="radio"
                      name="remont_type"
                      value="дизайнерский"
                      checked={form.remont_type === "дизайнерский"}
                      onChange={handleChange}
                    />
                    <span className={styles.radioText}>дизайнерский</span>
                  </label>

                  <label className={styles.customRadio}>
                    <input
                      type="radio"
                      name="remont_type"
                      value="требуется"
                      checked={form.remont_type === "требуется"}
                      onChange={handleChange}
                    />
                    <span className={styles.radioText}>требуется</span>
                  </label>
                </div>
              </div>
            </div>
          </div>

          <div className={styles.section}>
            <h3>Дополнительно</h3>
            <label className={styles.checkbox}>
              <input
                type="checkbox"
                name="elevator"
                checked={form.elevator}
                onChange={handleChange}
              />
              Лифт
            </label>
            <label className={styles.checkbox}>
              <input
                type="checkbox"
                name="garbage_chute"
                checked={form.garbage_chute}
                onChange={handleChange}
              />
              Мусоропровод
            </label>
            <label className={styles.checkbox}>
              <input
                type="checkbox"
                name="concierge"
                checked={form.concierge}
                onChange={handleChange}
              />
              Консьерж
            </label>
          </div>

          <PhotoUploader
            previewUrls={previewUrls}
            setPreviewUrls={setPreviewUrls}
            photos={photos}
            setPhotos={setPhotos}
          />

          <button type="submit" className={styles.button}>
            Добавить квартиру
          </button>
        </form>
      </div>
    </>
  );
};

export default AddApartmentPage;
