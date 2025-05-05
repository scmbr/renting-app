package dto

import "time"

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
	ID             uint      `json:"id"`
	UserID         uint      `json:"user_id"`
	ApartmentID    uint      `json:"apartment_id"`
	Title          string    `json:"title"`
	Pets           bool      `json:"pets"`
	Babies         bool      `json:"babies"`
	Smoking        bool      `json:"smoking"`
	Internet       bool      `json:"internet"`
	WashingMachine bool      `json:"washing_machine"`
	TV             bool      `json:"tv"`
	Conditioner    bool      `json:"conditioner"`
	Dishwasher     bool      `json:"dishwasher"`
	Concierge      bool      `json:"concierge"`
	Rent           float64   `json:"rent"`
	Deposit        float64   `json:"deposit"`
	RentalType     string    `json:"rental_type"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
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
