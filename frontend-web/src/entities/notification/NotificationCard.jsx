import { Link } from "react-router-dom";
import styles from "./NotificationCard.module.css";

export function NotificationCard({ notification }) {
  return (
    <li className={styles.card}>
      <Link to={`/advert/${notification.advert_id}`} className={styles.link}>
        <div className={styles.title}>{notification.title}</div>
        <div className={styles.content}>{notification.content}</div>
        <div className={styles.date}>
          {new Date(notification.created_at).toLocaleString()}
        </div>
      </Link>
    </li>
  );
}
