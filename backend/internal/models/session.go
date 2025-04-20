package models

import "time"

type Session struct {
	ID           int `gorm:"primaryKey"`
	UserID       int
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
	Browser      string
	OS           string
	IP           string
}
