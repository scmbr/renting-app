package server

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/scmbr/renting-app/internal/dto"
)

type WebSocketHub struct {
	clients map[uint]*websocket.Conn 
	mu      sync.Mutex
}

func NewWebSocketHub() *WebSocketHub {
	return &WebSocketHub{
		clients: make(map[uint]*websocket.Conn),
	}
}

func (hub *WebSocketHub) AddClient(userID uint, conn *websocket.Conn) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	hub.clients[userID] = conn
}

func (hub *WebSocketHub) RemoveClient(userID uint) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	if conn, ok := hub.clients[userID]; ok {
		conn.Close()
		delete(hub.clients, userID)
	}
}

func (hub *WebSocketHub) SendNotification(userID uint, notification dto.NotificationResponseDTO) error {
	hub.mu.Lock()
	conn, ok := hub.clients[userID]
	hub.mu.Unlock()

	if !ok {
		return fmt.Errorf("user %d not connected", userID)
	}

	message, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.TextMessage, message)
}
