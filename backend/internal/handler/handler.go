package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/scmbr/renting-app/docs"
	"github.com/scmbr/renting-app/internal/service"
	"github.com/scmbr/renting-app/pkg/auth"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh", h.refreshTokens)
		auth.POST("/verify", h.userVerify)
		auth.POST("/verify/resend", h.userVerifyResend)
		auth.POST("/forgot-password", h.forgotPass)
		auth.POST("/reset-password", h.resetPass)
	}
	authAuthorized := router.Group("/auth", h.userIdentity)
	{
		authAuthorized.POST("/logout", h.logOut)
		//authAuthorized.POST("/change-password", h.changePassword)
	}
	authenticated := router.Group("/", h.userIdentity)
	{
		authenticated.GET("/me", h.getCurrentUser)
		authenticated.POST("/upload-avatar", h.UploadAvatarHandler)
		apartment := authenticated.Group("/apartment")
		{
			apartment.GET("/", h.getAllApartments)
			apartment.GET("/:id", h.getApartmentById)
			apartment.POST("/", h.createApartment)
			apartment.DELETE("/:id", h.deleteApartment)
			apartment.PATCH("/:id", h.updateApartment)
		}
		advert := authenticated.Group("/advert")
		{
			advert.GET("/:id", h.getAdvertById)
			advert.GET("/", h.getAllAdverts)
			advert.POST("/", h.createAdvert)
			advert.DELETE("/:id", h.deleteAdvert)
			advert.PATCH("/:id", h.updateAdvert)
		}
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
