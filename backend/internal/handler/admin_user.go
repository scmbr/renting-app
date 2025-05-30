package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/renting-app/internal/dto"
)

func (h *Handler) adminGetAllUsers(c *gin.Context) {

	users, err := h.services.User.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

func (h *Handler) adminGetUserById(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.User.GetUserById(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) adminDeleteUserById(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.User.DeleteUserById(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) adminUpdateUserById(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input dto.UpdateUserAdmin
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.GetUserById(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Surname != nil {
		user.Surname = *input.Surname
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Birthdate != nil {
		user.Birthdate = *input.Birthdate
	}
	if input.Role != nil {
		user.Role = *input.Role
	}
	if input.IsActive != nil {
		user.IsActive = *input.IsActive
	}
	user, err = h.services.User.UpdateUserById(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
