package dto

import "github.com/scmbr/renting-app/internal/domain"

type CreateRoomRequest struct {
	ApartmentID uint    `json:"apartment_id" binding:"required"`
	Area        float64 `json:"area" binding:"required"`
}
type GetRoomResponse struct {
	ID          uint     `json:"id"`
	ApartmentID uint     `json:"apartment_id"`
	User        *GetUser `json:"user,omitempty"`
	Area        float64  `json:"area"`
	Status      string   `json:"status"`
}

func FromRoom(r *domain.Room) *GetRoomResponse {
	return &GetRoomResponse{
		ID:          r.ID,
		ApartmentID: r.ApartmentID,
		User:        FromUser(&r.Apartment.User),
		Area:        r.Area,
		Status:      r.Status,
	}
}
