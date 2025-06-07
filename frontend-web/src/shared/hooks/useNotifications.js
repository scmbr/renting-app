import { useEffect } from "react"
import { initNotificationSocket, closeNotificationSocket } from "@/shared/lib/websocket/notifications"
import { useNotificationStore } from "@/shared/store/notifications"

export function useNotifications(token) {
  const addNotification = useNotificationStore(state => state.addNotification)
  useEffect(() => {
    if (!token) return

    initNotificationSocket(token, (data) => {
      addNotification(data)
    })

    return () => {
      closeNotificationSocket()
    }
  }, [token])
}
