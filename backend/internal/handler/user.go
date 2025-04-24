package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"github.com/scmbr/renting-app/internal/dto"
	"github.com/sirupsen/logrus"
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

func (h *Handler) signUp(c *gin.Context) {
	var input dto.CreateUser
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.User.SignUp(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusCreated)
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type tokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userAgent := c.Request.Header.Get("User-Agent")
	ip := c.GetHeader("X-Forwarded-For")
	if ip == "" {
		ip = c.Request.RemoteAddr
	}

	ua := user_agent.New(userAgent)
	os := ua.OS()
	browser, _ := ua.Browser()
	res, err := h.services.User.SignIn(c, input.Email, input.Password, ip, os, browser)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

type refreshInput struct {
	Token string `json:"token" binding:"required"`
}

func (h *Handler) refreshTokens(c *gin.Context) {
	var input refreshInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	ip := c.GetHeader("X-Forwarded-For")
	if ip == "" {
		ip = c.Request.RemoteAddr
	}
	userAgent := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	os := ua.OS()
	browser, _ := ua.Browser()

	tokens, err := h.services.Session.RefreshSession(c.Request.Context(), input.Token, ip, os, browser)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

type response struct {
	Message string `json:"message"`
}
type VerifyRequest struct {
	Code string `json:"code"`
}

func (h *Handler) userVerify(c *gin.Context) {
	var input VerifyRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}
	if input.Code == "" {
		logrus.Error("code is empty")
		c.AbortWithStatusJSON(http.StatusBadRequest, response{"code is empty"})
		return
	}

	if err := h.services.User.VerifyEmail(c.Request.Context(), input.Code); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	c.JSON(http.StatusOK, response{"success"})
}
