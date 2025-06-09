package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
)

func (h *Handler) UploadAvatarHandler(c *gin.Context) {
	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID, _ := c.Get("userId")

	// Получение URL аватара через сервис загрузки
	avatarURL, err := h.services.User.UploadAvatarToS3(c.Request.Context(), fileHeader)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.User.UpdateAvatar(userID.(int), avatarURL)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{"avatar_url": avatarURL})
}
func (h *Handler) getCurrentUser(c *gin.Context) {
	userID, _ := c.Get("userId")

	user, err := h.services.User.GetUserById(userID.(int))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type tokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"Invalid input data"`
}
type response struct {
	Message string `json:"message"`
}
type VerifyRequest struct {
	Code string `json:"code"`
}

type ResetPasswordInput struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
type VerifyResendRequest struct {
	Email string `json:"email" example:"user@example.com"`
}

func (h *Handler) getUserById(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.User.GetUserById(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
func (h *Handler) updateCurrentUser(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		newErrorResponse(c, http.StatusUnauthorized, "user ID not found in context")
		return
	}

	var input dto.UpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	if err := h.services.User.UpdateMe(userID.(int), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to update user: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, response{Message: "user updated"})
}
