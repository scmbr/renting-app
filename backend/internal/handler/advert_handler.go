package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
)

// GET /advert/
func (h *Handler) getAllAdverts(c *gin.Context) {
	userId := c.MustGet("userId").(int)

	adverts, err := h.services.Advert.GetAllAdverts(c.Request.Context(), userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, adverts)
}

// GET /advert/:id
func (h *Handler) getAdvertById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id")
		return
	}

	userId := c.MustGet("userId").(int)

	advert, err := h.services.Advert.GetAdvertById(c.Request.Context(), userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, advert)
}

// POST /advert/
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

// DELETE /advert/:id
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

// PATCH /advert/:id
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
