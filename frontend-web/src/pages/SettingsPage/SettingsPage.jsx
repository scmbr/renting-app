import { useState } from "react";
import styles from "./SettingsPage.module.css";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";

const TABS = ["Аккаунт", "Безопасность", "Уведомления"];

const SettingsPage = () => {
  const [activeTab, setActiveTab] = useState("Аккаунт");

  const renderContent = () => {
    switch (activeTab) {
      case "Аккаунт":
        return (
          <div className={styles.section}>
            <h2>Данные профиля</h2>
            <label>
              Имя:
              <input type="text" placeholder="Введите имя" />
            </label>
            <label>
              Email:
              <input type="email" placeholder="example@email.com" />
            </label>
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
