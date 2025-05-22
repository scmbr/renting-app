package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"github.com/scmbr/renting-app/internal/dto"
	"github.com/sirupsen/logrus"
)

// @Summary      Регистрация пользователя
// @Description  Создаёт нового пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      dto.CreateUser  true  "Данные пользователя"
// @Success      201    {string}  string          "Created"
// @Failure      400    {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Router       /auth/sign-up [post]
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

// @Summary      Авторизация пользователя
// @Description  Вход пользователя в систему с возвратом access и refresh токенов
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      signInInput   true  "Данные для входа"
// @Success      200    {object}  tokenResponse
// @Failure      400    {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Router       /auth/sign-in [post]
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
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.SetCookie("refreshToken", res.RefreshToken, int(h.refreshTokenTTL.Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, tokenResponse{
		AccessToken: res.AccessToken,
	})
}

// @Summary      Обновление токенов
// @Description  Обновляет access и refresh токены по действующему refresh токену
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      refreshInput   true  "Refresh токен"
// @Success      200    {object}  tokenResponse
// @Failure      400    {object}  ErrorResponse
// @Failure      401    {object}  ErrorResponse
// @Router       /auth/refresh [post]
func (h *Handler) refreshTokens(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "no refresh token")
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
	claims, err := h.tokenManager.Parse(refreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "no refresh token")
		return
	}
	role := claims.Role
	tokens, err := h.services.Session.RefreshSession(c.Request.Context(), role, refreshToken, ip, os, browser)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.SetCookie("refreshToken", tokens.RefreshToken, int(h.refreshTokenTTL.Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, tokenResponse{
		AccessToken: tokens.AccessToken,
	})
}

// @Summary      Подтверждение электронной почты
// @Description  Проверяет код для верификации электронной почты пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      VerifyRequest   true  "Код для подтверждения"
// @Success      200    {object}  response
// @Failure      400    {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Router       /auth/verify [post]
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

// @Summary      Переотправка кода верификации
// @Description  Переотправялет код верификации
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      VerifyResendRequest   true  "Email для переотправки кода"
// @Success      200    {object}  response
// @Failure      400    {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Router       /auth/verify/resend [post]
func (h *Handler) userVerifyResend(c *gin.Context) {
	var input VerifyResendRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	err := h.services.User.ResendVerificationCode(c.Request.Context(), input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{"verification code resent successfully"})
}

// @Summary      Переотправка кода верификации
// @Description  Переотправялет код верификации
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      VerifyResendRequest   true  "Email для переотправки кода"
// @Success      200    {object}  response
// @Failure      400    {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Router       /auth/verify/resend [post]
func (h *Handler) forgotPass(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	err := h.services.User.ForgotPassword(c.Request.Context(), input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{"check your email for reset link"})
}

// resetPass обрабатывает запрос на сброс пароля через токен
// @Summary Reset Password
// @Tags auth
// @Description Сброс пароля пользователя через токен
// @Accept json
// @Produce json
// @Param input body ResetPasswordInput true "Данные для сброса пароля"
// @Success 200 {object} response
// @Failure 400,401,500 {object} response
// @Router /auth/reset-password [post]
func (h *Handler) resetPass(c *gin.Context) {
	var input ResetPasswordInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	if input.Token == "" || input.NewPassword == "" {
		newErrorResponse(c, http.StatusBadRequest, "token and new password must be provided")
		return
	}

	err := h.services.User.ResetPassword(c.Request.Context(), input.Token, input.NewPassword)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{"password reset successfully"})
}

func (h *Handler) logOut(c *gin.Context) {
	ip := c.GetHeader("X-Forwarded-For")
	if ip == "" {
		ip = c.Request.RemoteAddr
	}
	userAgent := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	os := ua.OS()
	browser, _ := ua.Browser()
	userId, _ := c.Get("userId")
	err := h.services.User.LogOut(c.Request.Context(), userId.(int), ip, os, browser)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.SetCookie("refreshToken", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, response{"user log out successfully"})
}
