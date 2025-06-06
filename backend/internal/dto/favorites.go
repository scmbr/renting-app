package dto

import "time"

type AddFavoriteDTO struct {
	AdvertID int `json:"advert_id" binding:"required"`
}

type FavoriteResponseDTO struct {
	ID        uint      `json:"id"`
	AdvertID  uint      `json:"advert_id" `
	CreatedAt time.Time `json:"created_at"`
}
