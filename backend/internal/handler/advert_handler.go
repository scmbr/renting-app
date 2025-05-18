package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
)

// @Summary      Получить все объявления пользователя
// @Tags         adverts
// @Security     ApiKeyAuth
// @Produce      json
// @Success      200 {array} dto.GetAdvertResponse
// @Failure      500 {object} ErrorResponse
// @Router       /advert [get]
func (h *Handler) getAllAdverts(c *gin.Context) {
	userId := c.MustGet("userId").(int)

	adverts, err := h.services.Advert.GetAllAdverts(c.Request.Context(), userId)
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
