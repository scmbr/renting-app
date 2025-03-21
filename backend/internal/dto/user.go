package dto

import "time"

type GetUser struct {
	Id        int       `json:"-"`
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surname" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Birthdate time.Time `json:"birthdate" binding:"required"`
	Role      string    `json:"role" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

type CreateUser struct {
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surname" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Birthdate time.Time `json:"birthdate" binding:"required"`
	Role      string    `json:"role" binding:"required"`
}
