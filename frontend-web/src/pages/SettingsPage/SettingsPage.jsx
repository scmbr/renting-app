import { useEffect, useState, useRef } from "react";
import styles from "./SettingsPage.module.css";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";
import api from "@/shared/api/axios";

const TABS = ["Аккаунт", "Безопасность", "Уведомления"];

const SettingsPage = () => {
  const [activeTab, setActiveTab] = useState("Аккаунт");
  const [userData, setUserData] = useState(null);
  const [formData, setFormData] = useState({
    name: "",
    surname: "",
    email: "",
    phone: "",
    city: "",
  });

  const fileInputRef = useRef(null);

  const fetchUserData = async () => {
    try {
      const response = await api.get("/me");
      setUserData(response.data);
    } catch (error) {
      console.error("Ошибка загрузки данных пользователя:", error);
    }
  };

  useEffect(() => {
    fetchUserData();
  }, []);

  useEffect(() => {
    if (userData) {
      setFormData({
        name: userData.name || "",
        surname: userData.surname || "",
        email: userData.email || "",
        phone: userData.phone || "",
        city: userData.city || "",
      });
    }
  }, [userData]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSave = async () => {
    try {
      await api.put("/me", formData);
      await fetchUserData();

      localStorage.setItem("city", formData.city);
      localStorage.setItem("name", formData.name);
      localStorage.setItem("surname", formData.surname);
      window.location.reload();
    } catch (error) {
      console.error("Ошибка при обновлении данных:", error);
      alert("Ошибка при обновлении данных");
    }
  };

  const handleAvatarClick = () => {
    fileInputRef.current.click();
  };

  const handleAvatarChange = async (e) => {
    const file = e.target.files[0];
    if (!file) return;

    try {
      const data = new FormData();
      data.append("avatar", file);

      const res = await api.post("/upload-avatar", data, {
        headers: { "Content-Type": "multipart/form-data" },
      });

      await fetchUserData();
      localStorage.setItem("avatarUrl", res.data.avatar_url);
      window.location.reload();
    } catch (error) {
      console.error("Ошибка при обновлении аватара:", error);
      alert("Ошибка при обновлении аватара");
    }
  };

  const renderContent = () => {
    switch (activeTab) {
      case "Аккаунт":
        return (
          <div className={styles.section}>
            <h2>Данные профиля</h2>

            <div className={styles.avatarContainer}>
              <img
                src={
                  userData?.profile_picture ||
                  "https://storage.yandexcloud.net/profile-pictures/user.png"
                }
                alt="Аватар пользователя"
                className={styles.avatar}
              />
              <span
                className={styles.updateAvatarText}
                onClick={handleAvatarClick}
                role="button"
                tabIndex={0}
                onKeyDown={(e) => {
                  if (e.key === "Enter" || e.key === " ") handleAvatarClick();
                }}
              >
                Изменить фотографию
              </span>
              <input
                type="file"
                accept="image/*"
                ref={fileInputRef}
                onChange={handleAvatarChange}
                style={{ display: "none" }}
              />
            </div>

            <label>
              Имя:
              <input
                type="text"
                name="name"
                placeholder="Введите имя"
                value={formData.name}
                onChange={handleChange}
              />
            </label>

            <label>
              Фамилия:
              <input
                type="text"
                name="surname"
                placeholder="Введите фамилию"
                value={formData.surname}
                onChange={handleChange}
              />
            </label>

            <label>
              Email:
              <input
                type="email"
                name="email"
                placeholder="example@email.com"
                value={formData.email}
                onChange={handleChange}
              />
            </label>

            <label>
              Номер:
              <input
                type="text"
                name="phone"
                placeholder="79999999999"
                value={formData.phone}
                onChange={handleChange}
              />
            </label>

            <label>
              Город:
              <input
                type="text"
                name="city"
                placeholder="Москва"
                value={formData.city}
                onChange={handleChange}
              />
            </label>

            <button onClick={handleSave} className={styles.saveButton}>
              Сохранить изменения
            </button>
          </div>
        );

      case "Безопасность":
        return (
          <div className={styles.section}>
            <h2>Изменение пароля</h2>
            <label>
              Новый пароль:
              <input type="password" placeholder="Новый пароль" />
            </label>
            <label>
              Подтверждение:
              <input type="password" placeholder="Подтвердите пароль" />
            </label>
          </div>
        );

      case "Уведомления":
        return (
          <div className={styles.section}>
            <h2>Настройки уведомлений</h2>
            <label>
              <input type="checkbox" />
              Получать email-уведомления
            </label>
            <label>
              <input type="checkbox" />
              Получать push-уведомления
            </label>
          </div>
        );

      default:
        return null;
    }
  };

  return (
    <>
      <SubNavbar />
      <NavPanel />
      <div className={styles.wrapper}>
        <h1 className={styles.title}>Настройки</h1>
        <div className={styles.settingsLayout}>
          <div className={styles.sidebar}>
            {TABS.map((tab) => (
              <button
                key={tab}
                className={`${styles.tabButton} ${
                  activeTab === tab ? styles.active : ""
                }`}
                onClick={() => setActiveTab(tab)}
              >
                {tab}
              </button>
            ))}
          </div>
          <div className={styles.content}>{renderContent()}</div>
        </div>
      </div>
    </>
  );
};

export default SettingsPage;
