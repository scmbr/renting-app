package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Surname        string
	Email          string `gorm:"unique;not null"`
	PasswordHash   string `gorm:"not null"`
	Birthdate      time.Time
	Role           string `gorm:"default:user;not null"`
	ProfilePicture string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Verified       bool
	Rating         bool
	Gender         bool
	IsActive       bool
}
