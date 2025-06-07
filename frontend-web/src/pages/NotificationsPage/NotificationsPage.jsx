import { useEffect } from "react";
import api from "@/shared/api/axios";
import { useNotificationStore } from "@/stores/useNotificationStore";
import {
  initNotificationSocket,
  closeNotificationSocket,
} from "@/shared/lib/websocket/notifications";
import { NotificationList } from "@/features/notifications/ui/NotificationList";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";
const NotificationsPage = () => {
  const addNotification = useNotificationStore(
    (state) => state.addNotification
  );
  const setNotifications = useNotificationStore(
    (state) => state.setNotifications
  );

  useEffect(() => {
    api
      .get("/notifications")
      .then((res) => {
        setNotifications(res.data);
      })
      .catch((err) => {
        console.error("Ошибка загрузки уведомлений", err);
      });

    const token = localStorage.getItem("accessToken");
    if (!token) return;

    initNotificationSocket(token, (data) => {
      addNotification(data);
    });

    return () => {
      closeNotificationSocket();
    };
  }, []);

  return (
    <>
      <SubNavbar />
      <NavPanel />
      <div className="p-4">
        <NotificationList />
      </div>
    </>
  );
};

export default NotificationsPage;
