package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vasya/renting-app/internal/cloud"
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
        newErrorResponse(c,http.StatusInternalServerError,err.Error())
        return
    }
    

    userID, _ := c.Get("userId")
  
    fmt.Printf("Current user ID: %d\n", userID)
    
    // Получение URL аватара через сервис загрузки
    avatarURL, err := cloud.UploadAvatarToS3(fileHeader)
    if err != nil {
        newErrorResponse(c,http.StatusInternalServerError,err.Error())
        return
    }

   
    err = h.services.Users.UpdateAvatar(userID.(int), avatarURL)
    if err != nil {
        newErrorResponse(c,http.StatusInternalServerError,err.Error())
        return
    }

    c.JSON(200, gin.H{"avatar_url": avatarURL})
}
