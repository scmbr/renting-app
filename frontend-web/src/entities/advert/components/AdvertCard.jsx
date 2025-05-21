import { Link } from "react-router-dom";
import styles from "./AdvertCard.module.css";

const AdvertCard = ({ advert }) => {
  const { id, title, rent, apartment } = advert;

  return (
    <Link to={`/advert/${id}`} className={styles.advertLink}>
      <div className={styles.advertCard}>
        <h3 className="text-lg font-bold">{title}</h3>
        <p>
          {apartment.city}, {apartment.street} {apartment.house}
        </p>
        <p>
          Этаж: {apartment.floor} | Комнат: {apartment.rooms}
        </p>
        <p className="font-semibold">{rent.toLocaleString()} ₽/мес</p>
      </div>
    </Link>
  );
};

export default AdvertCard;
