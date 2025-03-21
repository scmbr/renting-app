package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vasya/renting-app/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/check", h.userIdentity)
	}
	api := router.Group("/api", h.userIdentity)
	{
		api.POST("/users/upload-avatar", h.UploadAvatarHandler)

	}
	admin := router.Group("/admin", h.adminMiddleware)
	{
		admin.DELETE("/users/:id", h.deleteUserById)
		admin.GET("/users", h.getAllUsers)
		admin.GET("/users/:id", h.getUserById)

	}
	return router
}
