package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UploadAvatarHandler(c *gin.Context) {
	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID, _ := c.Get("userId")

	// Получение URL аватара через сервис загрузки
	avatarURL, err := h.services.Users.UploadAvatarToS3(c.Request.Context(), fileHeader)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Users.UpdateAvatar(userID.(int), avatarURL)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{"avatar_url": avatarURL})
}
func (h *Handler) getCurrentUser(c *gin.Context) {
	userID, _ := c.Get("userId")

	user, err := h.services.Users.GetUserById(userID.(int))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
