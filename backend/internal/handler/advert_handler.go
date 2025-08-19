package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
)

type AdvertListResponse struct {
	Total   int64                    `json:"total"`
	Adverts []*dto.GetAdvertResponse `json:"adverts"`
}

// @Summary      Получить все объявления пользователя
// @Tags         adverts
// @Security     ApiKeyAuth
// @Produce      json
// @Success      200 {array} dto.GetAdvertResponse
// @Failure      500 {object} ErrorResponse
// @Router       /advert [get]
func (h *Handler) getAllUserAdverts(c *gin.Context) {
	userId := c.MustGet("userId").(int)

	adverts, err := h.services.Advert.GetAllUserAdverts(c.Request.Context(), userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, adverts)
}

// @Summary      Получить объявление по ID
// @Tags         adverts
// @Security     ApiKeyAuth
// @Produce      json
// @Param        id   path      int  true  "Advert ID"
// @Success      200  {object}  dto.GetAdvertResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /advert/{id} [get]
func (h *Handler) getUserAdvertById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id")
		return
	}

	userId := c.MustGet("userId").(int)

	advert, err := h.services.Advert.GetUserAdvertById(c.Request.Context(), userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, advert)
}

// @Summary      Создать объявление
// @Tags         adverts
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        input  body      dto.CreateAdvertInput  true  "Advert input"
// @Success      201
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /advert [post]
func (h *Handler) createAdvert(c *gin.Context) {
	userId := c.MustGet("userId").(int)

	var input dto.CreateAdvertInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Advert.CreateAdvert(c.Request.Context(), userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusCreated)
}

// @Summary      Удалить объявление
// @Tags         adverts
// @Security     ApiKeyAuth
// @Produce      json
// @Param        id   path      int  true  "Advert ID"
// @Success      200
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /advert/{id} [delete]
func (h *Handler) deleteAdvert(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id")
		return
	}

	userId := c.MustGet("userId").(int)

	err = h.services.Advert.DeleteAdvert(c.Request.Context(), userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

// @Summary      Обновить объявление
// @Tags         adverts
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        id     path      int                   true  "Advert ID"
// @Param        input  body      dto.UpdateAdvertInput true  "Update input"
// @Success      200
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /advert/{id} [patch]
func (h *Handler) updateAdvert(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id")
		return
	}

	userId := c.MustGet("userId").(int)

	var input dto.UpdateAdvertInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Advert.UpdateAdvert(c.Request.Context(), userId, id, &input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
func (h *Handler) getAllAdverts(c *gin.Context) {
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
	userIdVal, exists := c.Get("userId")
	var userId *int
	if exists {
		uid := userIdVal.(int)
		userId = &uid
	} else {
		userId = nil
	}
	adverts, total, err := h.services.Advert.GetAllAdverts(c, userId, &filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, AdvertListResponse{
		Total:   total,
		Adverts: adverts,
	})
}
func (h *Handler) getAdvertById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid advert id"})
		return
	}

	advert, err := h.services.Advert.GetAdvertById(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, advert)
}

func parseIntOrZero(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func parseFloat32OrZero(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}

func parseIntOrDefault(s string, def int) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return n
}

func parseBoolPointer(s string) *bool {
	if s == "" {
		return nil
	}
	b, _ := strconv.ParseBool(s)
	return &b
}
