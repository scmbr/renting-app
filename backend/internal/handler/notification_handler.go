package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	serverws "github.com/scmbr/renting-app/internal/server"
	"github.com/scmbr/renting-app/pkg/auth"
)

type WebSocketHandler struct {
	hub          *serverws.WebSocketHub
	tokenManager auth.TokenManager
}

func NewWebSocketHandler(hub *serverws.WebSocketHub, tokenManager auth.TokenManager) *WebSocketHandler {
	return &WebSocketHandler{hub: hub, tokenManager: tokenManager}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // настроить под CORS
}

func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, err := h.tokenManager.Parse(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId, err := strconv.Atoi(claims.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	h.hub.AddClient(uint(userId), conn)
}
func (h *Handler) getUserNotifications(c *gin.Context) {
	userIDVal, exists := c.Get("userId")
	if !exists {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	id, ok := userIDVal.(int)
	if !ok {

		newErrorResponse(c, http.StatusInternalServerError, "wrong type")
		return
	}
	userID := uint(id)

	notifications, err := h.services.Notification.GetUserNotifications(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to fetch notifications")
		return
	}

	c.JSON(http.StatusOK, notifications)
}
