package dto

import "time"

type CreateApartmentInput struct {
	City             string  `json:"city" binding:"required"`
	Street           string  `json:"street" binding:"required"`
	District         string  `json:"district"`
	House            string  `json:"house"`
	Building         string  `json:"building"`
	Floor            int     `json:"floor" binding:"required"`
	ApartmentNumber  string  `json:"apartment_number"`
	Longitude        float64 `json:"longitude" binding:"required"`
	Latitude         float64 `json:"latitude" binding:"required"`
	Rooms            int     `json:"rooms" binding:"required"`
	Elevator         bool    `json:"elevator"`
	GarbageChute     bool    `json:"garbage_chute"`
	BathroomType     string  `json:"bathroom_type"`
	Concierge        bool    `json:"concierge"`
	ConstructionYear int     `json:"construction_year" binding:"required"`
	ConstructionType string  `json:"construction_type"`
	Remont           string  `json:"remont"`
}
type GetApartmentResponse struct {
	ID               uint      `json:"id"`
	UserID           uint      `json:"user_id"`
	City             string    `json:"city"`
	Street           string    `json:"street"`
	District         string    `json:"district"`
	House            string    `json:"house"`
	Building         string    `json:"building"`
	Floor            int       `json:"floor"`
	ApartmentNumber  string    `json:"apartment_number"`
	Longitude        float64   `json:"longitude"`
	Latitude         float64   `json:"latitude"`
	Rooms            int       `json:"rooms"`
	Elevator         bool      `json:"elevator"`
	GarbageChute     bool      `json:"garbage_chute"`
	BathroomType     string    `json:"bathroom_type"`
	Concierge        bool      `json:"concierge"`
	ConstructionYear int       `json:"construction_year"`
	ConstructionType string    `json:"construction_type"`
	Remont           string    `json:"remont"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Rating           float32   `json:"rating"`
	Status           string    `json:"status"`
}

type UpdateApartmentInput struct {
	City             *string  `json:"city,omitempty"`
	Street           *string  `json:"street,omitempty"`
	District         *string  `json:"district,omitempty"`
	House            *string  `json:"house,omitempty"`
	Building         *string  `json:"building,omitempty"`
	Floor            *int     `json:"floor,omitempty"`
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
