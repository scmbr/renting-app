package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
)

func (h *Handler) getAllFavorites(c *gin.Context) {
	userId := c.GetInt("userId")

	favorites, err := h.services.Favorites.GetAllFavorites(c, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, favorites)
}

func (h *Handler) addToFavorites(c *gin.Context) {
	userId := c.GetInt("userId")

	var input dto.AddFavoriteDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Favorites.AddToFavorites(c, userId, input.AdvertID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) removeFavorite(c *gin.Context) {
	userId := c.GetInt("userId")

	advertId, err := strconv.Atoi(c.Param("advertId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id")
		return
	}

	err = h.services.Favorites.RemoveFromFavorites(c, userId, advertId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) isFavorite(c *gin.Context) {
	userId := c.GetInt("userId")

	advertId, err := strconv.Atoi(c.Param("advertId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id")
		return
	}

	isFav, err := h.services.Favorites.IsFavorite(c, userId, advertId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"is_favorite": isFav})
}
