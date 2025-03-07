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
