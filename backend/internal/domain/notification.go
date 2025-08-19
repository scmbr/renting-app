package domain

import "time"

type Notification struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Type      string
	AdvertID  uint
	Title     string
	Content   string
	IsSent    bool
	IsRead    bool
	CreatedAt time.Time
	SentAt    *time.Time
}
