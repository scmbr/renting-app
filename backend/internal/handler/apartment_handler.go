package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
)

// @Summary Получить все квартиры
// @Description Получить список всех квартир пользователя
// @Tags apartments
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} ErrorResponse
// @Router /apartment/ [get]
func (h *Handler) getAllUserApartments(c *gin.Context) {
	userId, _ := c.Get("userId")
	apartments, err := h.services.Apartment.GetAllApartments(c, userId.(int))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"apartments": apartments,
	})

}

// @Summary Получить квартиру по ID
// @Description Получить квартиру по ID
// @Tags apartments
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "Apartment ID"
// @Success 200 {object} dto.GetApartmentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /apartment/{id} [get]
func (h *Handler) getUserApartmentById(c *gin.Context) {
	id := c.Param("id")
	apartmentId, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, _ := c.Get("userId")
	apartment, err := h.services.Apartment.GetApartmentById(c, userId.(int), apartmentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, apartment)
}

// @Summary Создать квартиру
// @Description Создать новую квартиру
// @Tags apartments
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body dto.CreateApartmentInput true "Apartment input"
// @Success 201
// @Failure 500 {object} ErrorResponse
// @Router /apartment/ [post]
func (h *Handler) createApartment(c *gin.Context) {
	var input dto.CreateApartmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}
	userId, _ := c.Get("userId")
	err := h.services.Apartment.CreateApartment(c, userId.(int), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusCreated)
}

// @Summary Удалить квартиру
// @Description Удалить квартиру по ID
// @Tags apartments
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "Apartment ID"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /apartment/{id} [delete]
func (h *Handler) deleteApartment(c *gin.Context) {
	//TODO:удалить квартиру
	id := c.Param("id")
	apartmentId, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, _ := c.Get("userId")
	err = h.services.Apartment.DeleteApartment(c, userId.(int), apartmentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

// @Summary Обновить квартиру
// @Description Обновить данные по квартире
// @Tags apartments
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Apartment ID"
// @Param input body dto.UpdateApartmentInput true "Apartment update input"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /apartment/{id} [patch]
func (h *Handler) updateApartment(c *gin.Context) {
	idParam := c.Param("id")
	apartmentId, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid apartment id")
		return
	}

	userIdRaw, exists := c.Get("userId")
	if !exists {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	userId := userIdRaw.(int)

	var input dto.UpdateApartmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input: "+err.Error())
		return
	}

	err = h.services.Apartment.UpdateApartment(c.Request.Context(), userId, apartmentId, &input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
