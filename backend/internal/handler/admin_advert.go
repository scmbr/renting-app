package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
	"gorm.io/gorm"
)

func (h *Handler) adminGetAllAdverts(c *gin.Context) {
	adverts, err := h.services.Advert.GetAllAdvertsAdmin(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get adverts"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"adverts": adverts,
	})
}

func (h *Handler) adminGetAdvertById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid advert id"})
		return
	}

	advert, err := h.services.Advert.GetAdvertByIdAdmin(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "advert not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get advert"})
		}
		return
	}
	c.JSON(http.StatusOK, advert)
}

func (h *Handler) adminUpdateAdvert(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid advert id"})
		return
	}

	var input dto.UpdateAdvertInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return
	}

	err = h.services.Advert.UpdateAdvertAdmin(c.Request.Context(), id, &input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "advert not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update advert"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "advert updated successfully"})
}

func (h *Handler) adminDeleteAdvert(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid advert id"})
		return
	}

	err = h.services.Advert.DeleteAdvertAdmin(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "advert not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete advert"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "advert deleted successfully"})
}
