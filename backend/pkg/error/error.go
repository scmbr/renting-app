package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}
type Response struct {
	Error   string        `json:"error"`
	Details []ErrorDetail `json:"details,omitempty"`
}

const (
	ErrInvalidAdvertID       = "invalid advert id"
	ErrInvalidAdvertInput    = "invalid advert input"
	ErrInvalidUserID         = "invalid user id"
	ErrInvalidApartmentID    = "invalid apartment id"
	ErrInvalidApartmentInput = "invalid apartment input"
	ErrInternalServer        = "internal server error"
	ErrUnauthorized          = "user unauthorized"
	ErrForbidden             = "access forbidden"
	ErrBadRequest            = "bad request"
)

var Messages = map[string]string{
	ErrInvalidAdvertID:       "Некорректный ID объявления",
	ErrInvalidAdvertInput:    "Некорректные входные данные объявления",
	ErrInvalidUserID:         "Некорректный ID пользователя",
	ErrInvalidApartmentID:    "Некорректный ID квартиры",
	ErrInvalidApartmentInput: "Некорректные входные данные квартиры",
	ErrInternalServer:        "Внутренняя ошибка сервера",
	ErrUnauthorized:          "Пользователь не авторизован",
	ErrForbidden:             "Доступ запрещён",
	ErrBadRequest:            "Некорректный запрос",
}

func Send(c *gin.Context, statusCode int, message string) {
	logrus.WithFields(logrus.Fields{
		"status": statusCode,
		"path":   c.FullPath(),
		"method": c.Request.Method,
	}).Error(message)

	c.AbortWithStatusJSON(statusCode, Response{
		Error: Messages[message],
	})
}

func Internal(c *gin.Context, err error) {
	logrus.Error(err)
	Send(c, http.StatusInternalServerError, ErrInternalServer)
}
