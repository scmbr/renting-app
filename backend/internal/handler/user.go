package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllUsers(c *gin.Context) {
	
	
	users,err:= h.services.Users.GetAllUsers()
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return 
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"users":users,
	})
}

func (h *Handler) getUserById(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
	user,err:= h.services.Users.GetUserById(userID)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return 
	}
	c.JSON(http.StatusOK,user)
}
func (h *Handler) UploadAvatarHandler(c *gin.Context) {
    fileHeader, err := c.FormFile("avatar")
    if err != nil {
        c.AbortWithStatusJSON(400, gin.H{"error": "Invalid file"})
        return
    }

    // Получение URL аватара через сервис загрузки
    avatarURL, err := h.services.Users.UploadAvatarToS3(fileHeader)
    if err != nil {
        c.AbortWithStatusJSON(500, gin.H{"error": "Failed to upload avatar"})
        return
    }

    userID,err := h.services.Users.GetCurrentUserId(c) // Реализуйте получение ID пользователя
    err = h.services.Users.UpdateAvatar(userID, avatarURL)
    if err != nil {
        c.AbortWithStatusJSON(500, gin.H{"error": "Failed to update avatar"})
        return
    }

    c.JSON(200, gin.H{"avatar_url": avatarURL})
}