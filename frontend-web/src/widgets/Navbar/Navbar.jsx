import { NavLink, useNavigate } from "react-router-dom";
import CitySelector from "./CitySelector";
import styles from "./Navbar.module.css";
import { useRef, useState } from "react";
import { useUser } from "@/shared/contexts/UserContext";
import api from "@/shared/api/axios";

const Navbar = ({ selectedCity, onCitySelect }) => {
  const { user, logout } = useUser();
  const [menuOpen, setMenuOpen] = useState(false);
  const menuRef = useRef(null);
  const navigate = useNavigate();

  const handleLogout = async () => {
    try {
      await api.post("/auth/logout");
    } catch (err) {
      console.error("Ошибка при выходе:", err);
    }

    logout();
    navigate("/login");
  };

  return (
    <nav className={styles.navbar}>
      <div className={styles.central}>
        <NavLink to="/" className={styles.logoLink}>
          <img src="/images/logo.png" alt="Дом" className={styles.logo} />
        </NavLink>

        <CitySelector selectedCity={selectedCity} onSelect={onCitySelect} />

        <div className={styles.spacer}></div>
        <div className={styles.navbarItems}>
          <NavLink
            to="/advert/add"
            className={({ isActive }) =>
              isActive ? `${styles.iconLink} ${styles.active}` : styles.iconLink
            }
          >
            <img src="/icons/add.svg" alt="Добавить" className={styles.icon} />
          </NavLink>
          <NavLink
            to="/my/advert"
            className={({ isActive }) =>
              isActive ? `${styles.iconLink} ${styles.active}` : styles.iconLink
            }
          >
            <img src="/icons/adverts.svg" alt="Мои" className={styles.icon} />
          </NavLink>
          <NavLink
            to="/notifications"
            className={({ isActive }) =>
              isActive ? `${styles.iconLink} ${styles.active}` : styles.iconLink
            }
          >
            <img
              src="/icons/notifications.svg"
              alt="Уведомления"
              className={styles.icon}
            />
          </NavLink>
          <NavLink
            to="/favorites"
            className={({ isActive }) =>
              isActive ? `${styles.iconLink} ${styles.active}` : styles.iconLink
            }
          >
            <img
              src="/icons/favourites.svg"
              alt="Избранное"
              className={styles.icon}
            />
          </NavLink>
          {user ? (
            <div className={styles.userMenu} ref={menuRef}>
              <button
                className={styles.userButton}
                onClick={() => setMenuOpen((prev) => !prev)}
              >
                <img src={user.avatarUrl} alt={``} className={styles.avatar} />
                {user.name} {user.surname}
              </button>
              {menuOpen && (
                <div className={styles.dropdown}>
                  <button onClick={handleLogout} className={styles.logout}>
                    Выйти
                  </button>
                </div>
              )}
            </div>
          ) : (
            <NavLink
              to="/login"
              className={({ isActive }) =>
                isActive ? `${styles.link} ${styles.active}` : styles.link
              }
            >
              Войти
            </NavLink>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
