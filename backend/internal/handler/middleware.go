package handler

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.Next()
		return
	}
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	claims, err := h.tokenManager.Parse(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	userId, err := strconv.Atoi(claims.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Set("userId", userId)
	c.Set("Role", claims.Role)
	c.Next()
}
func (h *Handler) adminMiddleware(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.Next()
		return
	}
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		c.Abort()
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		c.Abort()
		return
	}

	claims, err := h.tokenManager.Parse(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	if claims.Role != "admin" {
		newErrorResponse(c, http.StatusForbidden, "Access denied: admin rights required")
		c.Abort()
		return
	}

	c.Set("userId", claims.UserID)
	c.Next()
}

func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", os.Getenv("URL"))
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS,PATCH")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}
