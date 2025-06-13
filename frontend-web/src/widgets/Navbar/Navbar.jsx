import { NavLink, useNavigate } from "react-router-dom";
import CitySelector from "./CitySelector";
import styles from "./Navbar.module.css";
import { useRef, useState, useEffect } from "react";
import { useUser } from "@/shared/contexts/UserContext";
import api from "@/shared/api/axios";
import { useCityStore } from "@/stores/useCityStore";
import { useFiltersStore } from "@/stores/useFiltersStore";
import { nameToSlug } from "@/shared/constants/cities";

const Navbar = () => {
  const { user, logout } = useUser();
  const [menuOpen, setMenuOpen] = useState(false);
  const menuRef = useRef(null);
  const navigate = useNavigate();
  const city = useCityStore((state) => state.city);
  const citySlug = nameToSlug(city || "Москва");
  const updateFilter = useFiltersStore((state) => state.updateFilter);

  useEffect(() => {
    if (!city) return;
    updateFilter("city", city);
    const slug = nameToSlug(city);
    navigate(`/${slug}`);
  }, [city]);

  // Закрытие меню при клике вне его
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (menuRef.current && !menuRef.current.contains(event.target)) {
        setMenuOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const handleLogout = async () => {
    try {
      await api.post("/auth/logout");
    } catch (err) {
      console.error("Ошибка при выходе:", err);
    }
    logout();
    navigate("/login");
  };

  const handleNavigate = (path) => {
    navigate(path);
    setMenuOpen(false);
  };

  return (
    <nav className={styles.navbar}>
      <div className={styles.central}>
        <NavLink to={`/${citySlug}`} className={styles.logoLink}>
          <img src="/images/logo.png" alt="Дом" className={styles.logo} />
        </NavLink>

        <CitySelector />

        <div className={styles.spacer}></div>
        <div className={styles.navbarItems}>
          <NavLink
            to="/my/advert/add"
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
                <img src={user.avatarUrl} alt="" className={styles.avatar} />
                {user.name} {user.surname}
              </button>
              {menuOpen && (
                <div className={styles.dropdown}>
                  <button onClick={() => handleNavigate("/my/advert")}>
                    Профиль
                  </button>
                  <button onClick={() => handleNavigate("/my/advert")}>
                    Мои объявления
                  </button>
                  <button onClick={() => handleNavigate("/my/apartment")}>
                    Мои квартиры
                  </button>
                  <button onClick={() => handleNavigate("/notifications")}>
                    Уведомления
                  </button>
                  <button onClick={() => handleNavigate("/favorites")}>
                    Избранное
                  </button>
                  <button onClick={() => handleNavigate("/settings")}>
                    Настройки
                  </button>
                  <button onClick={handleLogout}>Выйти</button>
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
