package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
)

func (h *Handler) getAllApartments(c *gin.Context) {
	//TODO:получить все квартиры
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
func (h *Handler) getApartmentById(c *gin.Context) {
	//TODO:получить квартиру по id
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
func (h *Handler) createApartment(c *gin.Context) {
	//TODO:создать квартиру
	var input dto.CreateApartmentInput
	userId, _ := c.Get("userId")
	err := h.services.Apartment.CreateApartment(c, userId.(int), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusCreated)
}
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
