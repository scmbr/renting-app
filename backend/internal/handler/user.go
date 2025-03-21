package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vasya/renting-app/internal/cloud"
)

func (h *Handler) UploadAvatarHandler(c *gin.Context) {
	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID, _ := c.Get("userId")

	fmt.Printf("Current user ID: %d\n", userID)

	// Получение URL аватара через сервис загрузки
	avatarURL, err := cloud.UploadAvatarToS3(fileHeader)
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
