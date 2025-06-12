package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/scmbr/renting-app/docs"
	"github.com/scmbr/renting-app/internal/service"
	"github.com/scmbr/renting-app/pkg/auth"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services        *service.Services
	tokenManager    auth.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	wsHandler       *WebSocketHandler
}

func NewHandler(
	services *service.Services,
	tokenManager auth.TokenManager,
	accessTTL, refreshTTL time.Duration,
	wsHandler *WebSocketHandler,
) *Handler {
	return &Handler{
		services:        services,
		tokenManager:    tokenManager,
		refreshTokenTTL: refreshTTL,
		accessTokenTTL:  accessTTL,
		wsHandler:       wsHandler,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(
		corsMiddleware,
		gin.Recovery(),
		gin.Logger(),
	)

	publicAdverts := router.Group("/adverts")
	{
		publicAdverts.GET("", h.getAllAdverts)
		publicAdverts.GET("/:id", h.getAdvertById)

	}
	publicUser := router.Group("/users")
	{
		publicUser.GET("/:id", h.getUserById)
		publicUser.GET("/:id/reviews", h.getUserReviews)
	}
	publicApartment := router.Group("/apartment")
	{
		publicPhoto := publicApartment.Group("/:id/photos")
		{
			publicPhoto.GET("", h.getAllPhotos)
			publicPhoto.GET("/:photoId", h.getPhotoById)
		}
	}
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
	}
	router.GET("/notifications/ws", h.wsHandler.HandleWebSocket)
	authenticated := router.Group("/", h.userIdentity)
	{
		authenticated.GET("/me", h.getCurrentUser)
		authenticated.PUT("/me", h.updateCurrentUser)
		authenticated.POST("/upload-avatar", h.UploadAvatarHandler)
		reviews := authenticated.Group("/reviews")
		{
			reviews.POST("", h.createReview)
			reviews.PUT("/:id", h.updateReview)
			reviews.DELETE("/:id", h.deleteReview)
			reviews.GET("/my", h.getMyReviews)
		}
		userApartment := authenticated.Group("/my/apartment")
		{
			userApartment.GET("", h.getAllUserApartments)
			userApartment.GET("/:id", h.getUserApartmentById)
			userApartment.POST("", h.createApartment)
			userApartment.DELETE("/:id", h.deleteApartment)
			userApartment.PATCH("/:id", h.updateApartment)
			photo := userApartment.Group("/:id/photos")
			{
				photo.GET("/:photoId", h.getPhotoById)
				photo.POST("/batch", h.addPhotos)
				photo.DELETE("/:photoId", h.deletePhoto)
				photo.PATCH("/:photoId/set-cover", h.setCover)
				photo.PATCH("/replace", h.replacePhotos)
			}
		}
		userAdvert := authenticated.Group("/my/advert")
		{
			userAdvert.GET("/:id", h.getUserAdvertById)
			userAdvert.GET("", h.getAllUserAdverts)
			userAdvert.POST("", h.createAdvert)
			userAdvert.DELETE("/:id", h.deleteAdvert)
			userAdvert.PUT("/:id", h.updateAdvert)
		}
		favorites := authenticated.Group("/favorites")
		{
			favorites.GET("", h.getAllFavorites)
			favorites.POST("", h.addToFavorites)
			favorites.DELETE("/:advertId", h.removeFavorite)
			favorites.GET("/:advertId/check", h.isFavorite)
		}

		authenticated.GET("/notifications", h.getUserNotifications)

	}
	admin := router.Group("/admin", h.adminMiddleware)
	{

		users := admin.Group("/users")
		{
			users.GET("/", h.adminGetAllUsers)
			users.GET("/:id", h.adminGetUserById)
			users.PUT("/:id", h.adminUpdateUserById)
			users.DELETE("/:id", h.adminDeleteUserById)
		}

		apartment := admin.Group("/apartments")
		{
			apartment.GET("/", h.adminGetAllApartments)
			apartment.GET("/:id", h.adminGetApartmentById)
			apartment.PUT("/:id", h.adminUpdateApartment)
			apartment.DELETE("/:id", h.adminDeleteApartment)
		}

		advert := admin.Group("/adverts")
		{
			advert.GET("/", h.adminGetAllAdverts)
			advert.GET("/:id", h.adminGetAdvertById)
			advert.PUT("/:id", h.adminUpdateAdvert)
			advert.DELETE("/:id", h.adminDeleteAdvert)
		}
	}
	return router
}
