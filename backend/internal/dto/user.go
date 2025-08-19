package dto

import (
	"time"

	"github.com/scmbr/renting-app/internal/domain"
)

type GetUser struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	City           string    `json:"city"`
	ProfilePicture string    `json:"profile_picture"`
	Birthdate      time.Time `json:"birthdate"`
	CreatedAt      time.Time `json:"created_at"`
	Role           string    `json:"role"`
	Rating         float64   `json:"rating"`
	Verified       bool      `json:"verified"`
	IsActive       bool      `json:"is_active"`
}

type CreateUser struct {
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surname" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=8"`
	Birthdate time.Time `json:"birthdate"`
	City      string    `json:"city"`
}

type UpdateUser struct {
	Name      *string    `json:"name"`
	Surname   *string    `json:"surname"`
	Email     *string    `json:"email"`
	Birthdate *time.Time `json:"birthdate"`
	City      *string    `json:"city"`
	Phone     *string    `json:"phone"`
}
type UpdateUserAdmin struct {
	Name      *string    `json:"name"`
	Surname   *string    `json:"surname"`
	Email     *string    `json:"email"`
	Birthdate *time.Time `json:"birthdate"`
	Role      *string    `json:"role"`
	IsActive  *bool      `json:"is_active"`
}

func FromUser(u *domain.User) *GetUser {
	return &GetUser{
		Id:             int(u.ID),
		Name:           u.Name,
		Surname:        u.Surname,
		Email:          u.Email,
		ProfilePicture: u.ProfilePicture,
		Birthdate:      u.Birthdate,
		Role:           u.Role,
		Verified:       u.Verified,
		IsActive:       u.IsActive,
	}
}
