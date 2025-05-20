import { Link } from "react-router-dom";
import CitySelector from "./CitySelector";

const Navbar = ({ selectedCity, onCitySelect }) => {
  return (
    <nav style={{
      padding: "1rem",
      backgroundColor: "#f3f3f3",
      display: "flex",
      alignItems: "center",
      gap: "1rem"
    }}>
      <Link to="/">Home</Link>
      <Link to="/about">About</Link>
      <Link to="/contact">Contact</Link>

      <CitySelector
        selectedCity={selectedCity}
        onSelect={onCitySelect}
      />
    </nav>
  );
};

export default Navbar;
