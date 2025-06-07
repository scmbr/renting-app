import styles from "./NotificationList.module.css";
import { useNotificationStore } from "@/stores/useNotificationStore";
import { NotificationCard } from "@/entities/notification/NotificationCard";

export function NotificationList() {
  const notifications = useNotificationStore((state) => state.notifications);

  return (
    <div className={styles.container}>
      <h2 className={styles.heading}>Уведомления</h2>
      <ul className={styles.list}>
        {notifications.map((n, i) => (
          <NotificationCard key={i} notification={n} />
        ))}
      </ul>
    </div>
  );
}
