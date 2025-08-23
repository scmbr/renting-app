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
	filter := dto.AdvertFilter{
		City:               c.Query("city"),
		District:           c.Query("district"),
		BathroomType:       c.Query("bathroom_type"),
		Remont:             c.Query("remont"),
		RentalType:         c.Query("rental_type"),
		Rooms:              parseIntOrZero(c.Query("rooms")),
		PriceMin:           parseIntOrZero(c.Query("price_from")),
		PriceMax:           parseIntOrZero(c.Query("price_to")),
		FloorMin:           parseIntOrZero(c.Query("floor_min")),
		FloorMax:           parseIntOrZero(c.Query("floor_max")),
		YearMin:            parseIntOrZero(c.Query("year_min")),
		YearMax:            parseIntOrZero(c.Query("year_max")),
		ApartmentRatingMin: parseFloat32OrZero(c.Query("rating_min")),
		LandlordRatingMin:  parseFloat32OrZero(c.Query("rating_min")),
		Limit:              parseIntOrDefault(c.Query("limit"), 20),
		Offset:             parseIntOrDefault(c.Query("offset"), 0),
		SortBy:             c.DefaultQuery("sort_by", "created_at"),
		Order:              c.DefaultQuery("order", "desc"),
		Lat:                parseFloat32OrZero(c.Query("lat")),
		Lng:                parseFloat32OrZero(c.Query("lng")),
	}

	filter.Elevator = parseBoolPointer(c.Query("elevator"))
	filter.Concierge = parseBoolPointer(c.Query("concierge"))
	filter.Pets = parseBoolPointer(c.Query("pets"))
	filter.Babies = parseBoolPointer(c.Query("babies"))
	filter.Smoking = parseBoolPointer(c.Query("smoking"))
	filter.Internet = parseBoolPointer(c.Query("internet"))
	filter.WashingMachine = parseBoolPointer(c.Query("washing_machine"))
	filter.TV = parseBoolPointer(c.Query("tv"))
	filter.Conditioner = parseBoolPointer(c.Query("conditioner"))
	filter.Dishwasher = parseBoolPointer(c.Query("dishwasher"))
	if filter.Limit <= 0 || filter.Limit > 100 {
		filter.Limit = 20
	}
	if filter.Offset < 0 {
		filter.Offset = 0
	}
	adverts, total, err := h.services.Advert.GetAllAdvertsAdmin(c.Request.Context(), &filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get adverts"})
		return
	}
	c.JSON(http.StatusOK, AdvertListResponse{
		Total:   total,
		Adverts: adverts,
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
