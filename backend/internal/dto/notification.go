package dto

import "time"

type NotificationDTO struct {
	UserID   uint   `json:"user_id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	AdvertId uint   `json:"advert_id"`
	Content  string `json:"content"`
}
type NotificationResponseDTO struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	AdvertId  uint      `json:"advert_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}
