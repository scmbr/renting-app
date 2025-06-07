import { NavLink } from "react-router-dom";
import styles from "./NavPanel.module.css";

const NavPanel = () => {
  return (
    <nav className={styles.subnavbar}>
      <div className={styles.central}>
        <NavLink
          to="/my/advert"
          className={({ isActive }) =>
            isActive ? `${styles.link} ${styles.active}` : styles.link
          }
        >
          Мои объявления
        </NavLink>
        <NavLink
          to="/my/apartment"
          className={({ isActive }) =>
            isActive ? `${styles.link} ${styles.active}` : styles.link
          }
        >
          Мои квартиры
        </NavLink>
        <NavLink
          to="/notifications"
          className={({ isActive }) =>
            isActive ? `${styles.link} ${styles.active}` : styles.link
          }
        >
          Уведомления
        </NavLink>
        <NavLink
          to="/favorites"
          className={({ isActive }) =>
            isActive ? `${styles.link} ${styles.active}` : styles.link
          }
        >
          Избранное
        </NavLink>
        <NavLink
          to="/settings"
          className={({ isActive }) =>
            isActive ? `${styles.link} ${styles.active}` : styles.link
          }
        >
          Настройки
        </NavLink>
      </div>
    </nav>
  );
};

export default NavPanel;
