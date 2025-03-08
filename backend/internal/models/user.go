package models

import "time"

type User struct {
	Id        int       `json:"-" db:"id"`
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surename" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	PasswordHash  string    `json:"password" binding:"required"`
	Birthdate time.Time `json:"birthdate" binding:"required"`
	Role      int       `json:"role" binding:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	IsActive  bool `json:"-"`
}


