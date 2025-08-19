package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string
	Surname          string
	Email            string `gorm:"unique;not null"`
	Phone            string
	PasswordHash     string `gorm:"not null"`
	Birthdate        time.Time
	Role             string `gorm:"default:user;not null"`
	ProfilePicture   string
	City             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	VerificationCode string `gorm:"column:verification_code"`
	Verified         bool
	Rating           float32
	Session          []Session `gorm:"foreignKey:UserID"`
	Gender           int       `gorm:"default:0;not null"`
	ResetToken       string
	IsActive         bool
}
