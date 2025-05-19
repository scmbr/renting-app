package dto

import (
	"time"

	"github.com/scmbr/renting-app/internal/models"
)

type CreateAdvertInput struct {
	ApartmentID    uint    `json:"apartment_id" binding:"required"`
	Title          string  `json:"title" binding:"required"`
	Pets           bool    `json:"pets"`
	Babies         bool    `json:"babies"`
	Smoking        bool    `json:"smoking"`
	Internet       bool    `json:"internet"`
	WashingMachine bool    `json:"washing_machine"`
	TV             bool    `json:"tv"`
	Conditioner    bool    `json:"conditioner"`
	Dishwasher     bool    `json:"dishwasher"`
	Concierge      bool    `json:"concierge"`
	Rent           float64 `json:"rent" binding:"required"`
	Deposit        float64 `json:"deposit" binding:"required"`
	RentalType     string  `json:"rental_type" binding:"required"`
}
type GetAdvertResponse struct {
	ID             uint                  `json:"id"`
	UserID         uint                  `json:"user_id"`
	ApartmentID    uint                  `json:"apartment_id"`
	Apartment      *GetApartmentResponse `json:"apartment"`
	Title          string                `json:"title"`
	Pets           bool                  `json:"pets"`
	Babies         bool                  `json:"babies"`
	Smoking        bool                  `json:"smoking"`
	Internet       bool                  `json:"internet"`
	WashingMachine bool                  `json:"washing_machine"`
	TV             bool                  `json:"tv"`
	Conditioner    bool                  `json:"conditioner"`
	Dishwasher     bool                  `json:"dishwasher"`
	Concierge      bool                  `json:"concierge"`
	Rent           float64               `json:"rent"`
	Deposit        float64               `json:"deposit"`
	RentalType     string                `json:"rental_type"`
	Status         string                `json:"status"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
}

type UpdateAdvertInput struct {
	Title          *string  `json:"title"`
	Pets           *bool    `json:"pets"`
	Babies         *bool    `json:"babies"`
	Smoking        *bool    `json:"smoking"`
	Internet       *bool    `json:"internet"`
	WashingMachine *bool    `json:"washing_machine"`
	TV             *bool    `json:"tv"`
	Conditioner    *bool    `json:"conditioner"`
	Dishwasher     *bool    `json:"dishwasher"`
	Concierge      *bool    `json:"concierge"`
	Rent           *float64 `json:"rent"`
	Deposit        *float64 `json:"deposit"`
	RentalType     *string  `json:"rental_type"`
	Status         *string  `json:"status"`
}
type AdvertFilter struct {
	City     string
	District string
	Rooms    int

	PriceMin int
	PriceMax int

	FloorMin int
	FloorMax int

	YearMin int
	YearMax int

	ApartmentRatingMin float32
	LandlordRatingMin  float32
	BathroomType       string
	Remont             string

	Elevator  *bool
	Concierge *bool

	Pets           *bool
	Babies         *bool
	Smoking        *bool
	Internet       *bool
	WashingMachine *bool
	TV             *bool
	Conditioner    *bool
	Dishwasher     *bool

	RentalType string

	Limit  int
	Offset int
	SortBy string
	Order  string
}

func FromAdvert(advert models.Advert) *GetAdvertResponse {
	return &GetAdvertResponse{
		ID:             advert.ID,
		UserID:         advert.UserID,
		ApartmentID:    advert.ApartmentID,
		Apartment:      FromApartment(&advert.Apartment),
		CreatedAt:      advert.CreatedAt,
		UpdatedAt:      advert.UpdatedAt,
		Status:         advert.Status,
		Title:          advert.Title,
		Pets:           advert.Pets,
		Babies:         advert.Babies,
		Smoking:        advert.Smoking,
		Internet:       advert.Internet,
		WashingMachine: advert.WashingMachine,
		TV:             advert.TV,
		Conditioner:    advert.Conditioner,
		Concierge:      advert.Concierge,
		Rent:           advert.Rent,
		Deposit:        advert.Deposit,
		RentalType:     advert.RentalType,
	}
}
