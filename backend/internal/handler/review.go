package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
)

func (h *Handler) createReview(c *gin.Context) {
	authorID := uint(c.GetInt("userId"))

	var input dto.CreateReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if authorID == input.TargetID {
		newErrorResponse(c, http.StatusBadRequest, "cannot review yourself")
		return
	}

	reviewResponse, err := h.services.Review.Create(c, authorID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, reviewResponse)
}

func (h *Handler) updateReview(c *gin.Context) {
	userID := uint(c.GetInt("userId"))

	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid review id")
		return
	}

	var input dto.UpdateReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	reviewResponse, err := h.services.Review.Update(c, userID, uint(reviewID), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, reviewResponse)
}

func (h *Handler) deleteReview(c *gin.Context) {
	userID := uint(c.GetInt("userId"))

	reviewID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid review id")
		return
	}

	err = h.services.Review.Delete(c, userID, uint(reviewID))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) getMyReviews(c *gin.Context) {
	userID := uint(c.GetInt("userId"))

	reviewsResponse, err := h.services.Review.GetByAuthorID(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, reviewsResponse)
}


func (h *Handler) getUserReviews(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	reviewsResponse, err := h.services.Review.GetByTargetID(c, uint(userID))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, reviewsResponse)
}