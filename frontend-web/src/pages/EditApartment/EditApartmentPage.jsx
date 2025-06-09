import React, { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import api from "@/shared/api/axios";
import styles from "./EditApartmentPage.module.css";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import { MapGLForm } from "@/widgets/Map/MapGLForm";
import AddressSuggester from "@/features/add-apartment/addressSuggester.jsx";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";
import PhotoUploader from "@/features/upload-photo/PhotoUploader";

const EditApartmentPage = () => {
  const { id } = useParams();
  const navigate = useNavigate();

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
    remont_type: "",
  });

  const [previewUrls, setPreviewUrls] = useState([]);
  const [photos, setPhotos] = useState([]);
  const [photosToDelete, setPhotosToDelete] = useState([]);
  const [error, setError] = useState(null);
  const [mapCenter, setMapCenter] = useState([37.620393, 55.75396]);
  const [mapMarker, setMapMarker] = useState([37.620393, 55.75396]);

  useEffect(() => {
    const fetchApartment = async () => {
      try {
        const res = await api.get(`/my/apartment/${id}`);
        const apt = res.data;

        setForm({
          city: apt.city || "",
          street: apt.street || "",
          building: apt.building || "",
          floor: apt.floor || "",
          longitude: apt.longitude,
          latitude: apt.latitude,
          rooms: apt.rooms,
          area: apt.area,
          elevator: apt.elevator,
          garbage_chute: apt.garbage_chute,
          bathroom_type: apt.bathroom_type,
          concierge: apt.concierge,
          construction_year: apt.construction_year,
          construction_type: apt.construction_type,
          remont_type: apt.remont_type,
        });

        setMapCenter([apt.longitude, apt.latitude]);
        setMapMarker([apt.longitude, apt.latitude]);

        const photoRes = await api.get(`/apartment/${id}/photos`);

        setPreviewUrls(
          photoRes.data.map((photo) => ({ id: photo.id, url: photo.url }))
        );
      } catch (err) {
        console.error(err);
        setError("Не удалось загрузить данные квартиры");
      }
    };

    fetchApartment();
  }, [id]);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
    }));
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
      ...form,
      floor: Number(form.floor),
      longitude: Number(form.longitude),
      latitude: Number(form.latitude),
      rooms: Number(form.rooms),
      area: form.area ? Number(form.area) : 0,
      construction_year: Number(form.construction_year),
    };

    try {
      await api.patch(`/my/apartment/${id}`, payload);

      await Promise.all(
        photosToDelete.map((photoId) =>
          api.delete(`/my/apartment/${id}/photos/${photoId}`)
        )
      );

      if (photos.length > 0) {
        const photoForm = new FormData();
        photos.forEach((file) => photoForm.append("photos", file));

        await api.post(`/my/apartment/${id}/photos`, photoForm, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });
      }

      navigate("/my/apartment");
    } catch (err) {
      console.error(err);
      setError("Ошибка при обновлении квартиры");
    }
  };

  return (
    <>
      <SubNavbar />
      <NavPanel />
      <div className={styles.wrapper}>
        <h2 className={styles.title}>Редактирование квартиры</h2>
        {error && <p className={styles.error}>{error}</p>}

        <form onSubmit={handleSubmit} className={styles.form}>
          <div className={styles.section}>
            <h3>Адрес</h3>
            <AddressSuggester
              className={styles.addressSuggester}
              location={`${form.longitude},${form.latitude}`}
              onSelect={handleAddressChange}
              value={`${form.street}, ${form.building}`}
            />

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
            setPhotosToDelete={setPhotosToDelete}
          />

          <button type="submit" className={styles.button}>
            Сохранить изменения
          </button>
        </form>
      </div>
    </>
  );
};

export default EditApartmentPage;
