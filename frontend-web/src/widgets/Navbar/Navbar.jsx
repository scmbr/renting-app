import { NavLink } from "react-router-dom";
import CitySelector from "./CitySelector";
import styles from "./Navbar.module.css";

const Navbar = ({ selectedCity, onCitySelect }) => {
  return (
    <nav className={styles.navbar}>
      <CitySelector selectedCity={selectedCity} onSelect={onCitySelect} />
      <div className={styles.spacer}></div>
      <NavLink
        to="/login"
        className={({ isActive }) =>
          isActive ? `${styles.link} ${styles.active}` : styles.link
        }
      >
        Войти
      </NavLink>
    </nav>
  );
};

export default Navbar;
