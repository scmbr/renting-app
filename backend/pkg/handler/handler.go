package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vasya/renting-app/pkg/service"
)

type Handler struct {
	services *service.Services
}
func NewHandler(services *service.Services) *Handler{
	return &Handler{services:services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth:=router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/check", h.userIdentity)
	}
	router.GET("/users", h.getAllUsers)
	router.GET("/users/:id", h.getUserById)
	return router
}