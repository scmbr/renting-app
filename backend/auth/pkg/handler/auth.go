package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vasya/auth"
)

func (h *Handler) signUp(c *gin.Context) {
	var input auth.User
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	id,err:= h.services.Authorization.CreateUser(input)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return 
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"id":id,
	})
}
type signInInput struct{
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	token,err:= h.services.Authorization.GenerateToken(input.Email,input.Password)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return 
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"token":token,
	})
}