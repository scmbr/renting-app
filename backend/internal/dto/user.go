package dto

import (
	"time"

	"github.com/scmbr/renting-app/internal/models"
)

type GetUser struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Surname          string    `json:"surname"`
	Email            string    `json:"email"`
	Birthdate        time.Time `json:"birthdate"`
	Role             string    `json:"role"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	VerificationCode string    `json:"verification_code"`
	Verified         bool      `json:"verified"`
	IsActive         bool      `json:"is_active"`
}

type CreateUser struct {
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surname" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=8"`
	Birthdate time.Time `json:"birthdate"`
}
type UpdateUserAdmin struct {
	Name      *string    `json:"name"`
	Surname   *string    `json:"surname"`
	Email     *string    `json:"email"`
	Birthdate *time.Time `json:"birthdate"`
	Role      *string    `json:"role"`
	IsActive  *bool      `json:"is_active"`
}

func FromUser(u *models.User) *GetUser {
	return &GetUser{
		Id:               int(u.ID),
		Name:             u.Name,
		Surname:          u.Surname,
		Email:            u.Email,
		Birthdate:        u.Birthdate,
		Role:             u.Role,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
		VerificationCode: u.VerificationCode,
		Verified:         u.Verified,
		IsActive:         u.IsActive,
	}
}
