package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/service"
	"github.com/scmbr/renting-app/pkg/auth"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh", h.refreshTokens)
	}
	client := router.Group("/client", h.userIdentity)
	{
		client.GET("/me", h.getCurrentUser)
		client.POST("/upload-avatar", h.UploadAvatarHandler)
	}
	admin := router.Group("/admin", h.adminMiddleware)
	{
		admin.DELETE("/users/:id", h.deleteUserById)
		admin.PUT("/users/:id", h.updateUserById)
		admin.GET("/users", h.getAllUsers)
		admin.GET("/users/:id", h.getUserById)

	}
	return router
}
