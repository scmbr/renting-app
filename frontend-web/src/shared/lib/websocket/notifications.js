let socket = null;

export function initNotificationSocket(token, onMessage) {
  const ws = new WebSocket(`ws://backend:8000/notifications/ws?token=${token}`);
  ws.onopen = () => {
    console.log("WebSocket подключен");
  };
  ws.onmessage = (event) => {
    try {
      const data = JSON.parse(event.data);
      onMessage?.(data);
    } catch (error) {
      console.error("Ошибка парсинга сообщения:", error);
    }
  };
  ws.onerror = (err) => {
    console.error("WebSocket ошибка:", err);
  };
  ws.onclose = () => {
    console.log("WebSocket отключён");
  };
  socket = ws;
}

export function closeNotificationSocket() {
  if (socket) {
    socket.close();
    socket = null;
  }
}
