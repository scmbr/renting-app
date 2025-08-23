package dto

import (
	"time"

	"github.com/scmbr/renting-app/internal/domain"
)

type CreateApartmentInput struct {
	City             string  `json:"city" binding:"required"`
	Street           string  `json:"street" binding:"required"`
	Building         string  `json:"building"`
	Floor            int     `json:"floor" binding:"required"`
	Longitude        float64 `json:"longitude" binding:"required"`
	Latitude         float64 `json:"latitude" binding:"required"`
	Rooms            int     `json:"rooms" binding:"required"`
	Area             int     `json:"area"`
	Elevator         bool    `json:"elevator"`
	GarbageChute     bool    `json:"garbage_chute"`
	BathroomType     string  `json:"bathroom_type"`
	Concierge        bool    `json:"concierge"`
	ConstructionYear int     `json:"construction_year" binding:"required"`
	ConstructionType string  `json:"construction_type"`
	Remont           string  `json:"remont"`
}
type GetApartmentResponse struct {
	ID               uint                        `json:"id"`
	UserID           uint                        `json:"user_id"`
	City             string                      `json:"city"`
	Street           string                      `json:"street"`
	District         string                      `json:"district"`
	House            string                      `json:"house"`
	Building         string                      `json:"building"`
	Floor            int                         `json:"floor"`
	ApartmentNumber  string                      `json:"apartment_number"`
	Longitude        float64                     `json:"longitude"`
	Latitude         float64                     `json:"latitude"`
	Rooms            int                         `json:"rooms"`
	Area             int                         `json:"area"`
	Elevator         bool                        `json:"elevator"`
	GarbageChute     bool                        `json:"garbage_chute"`
	BathroomType     string                      `json:"bathroom_type"`
	Concierge        bool                        `json:"concierge"`
	ConstructionYear int                         `json:"construction_year"`
	ConstructionType string                      `json:"construction_type"`
	Remont           string                      `json:"remont"`
	CreatedAt        time.Time                   `json:"created_at"`
	UpdatedAt        time.Time                   `json:"updated_at"`
	Rating           float32                     `json:"rating"`
	Status           string                      `json:"status"`
	ApartmentPhotos  []GetApartmentPhotoResponse `json:"apartment_photos"`
}

type UpdateApartmentInput struct {
	City             *string  `json:"city,omitempty"`
	Street           *string  `json:"street,omitempty"`
	District         *string  `json:"district,omitempty"`
	House            *string  `json:"house,omitempty"`
	Building         *string  `json:"building,omitempty"`
	Floor            *int     `json:"floor,omitempty"`
	Area             *int     `json:"area,omitempty"`
	ApartmentNumber  *string  `json:"apartment_number,omitempty"`
	Longitude        *float64 `json:"longitude,omitempty"`
	Latitude         *float64 `json:"latitude,omitempty"`
	Rooms            *int     `json:"rooms,omitempty"`
	Elevator         *bool    `json:"elevator,omitempty"`
	GarbageChute     *bool    `json:"garbage_chute,omitempty"`
	BathroomType     *string  `json:"bathroom_type,omitempty"`
	Concierge        *bool    `json:"concierge,omitempty"`
	ConstructionYear *int     `json:"construction_year,omitempty"`
	ConstructionType *string  `json:"construction_type,omitempty"`
	Remont           *string  `json:"remont,omitempty"`
	Status           *string  `json:"status,omitempty"`
}

func FromApartment(apartment *domain.Apartment) *GetApartmentResponse {
	if apartment == nil {
		return nil
	}

	return &GetApartmentResponse{
		ID:               apartment.ID,
		UserID:           apartment.UserID,
		City:             apartment.City,
		Street:           apartment.Street,
		Building:         apartment.Building,
		Floor:            apartment.Floor,
		Area:             apartment.Area,
		ApartmentNumber:  apartment.ApartmentNumber,
		Longitude:        apartment.Longitude,
		Latitude:         apartment.Latitude,
		Rooms:            apartment.Rooms,
		Elevator:         apartment.Elevator,
		GarbageChute:     apartment.GarbageChute,
		BathroomType:     apartment.BathroomType,
		Concierge:        apartment.Concierge,
		ConstructionYear: apartment.ConstructionYear,
		ConstructionType: apartment.ConstructionType,
		Remont:           apartment.Remont,
		CreatedAt:        apartment.CreatedAt,
		UpdatedAt:        apartment.UpdatedAt,
		Rating:           apartment.Rating,
		Status:           apartment.Status,
	}
}
