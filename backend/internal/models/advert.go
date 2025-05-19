package models

import (
	"time"
)

type Advert struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null;index"`
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ApartmentID uint      `gorm:"not null;index"`
	Apartment   Apartment `gorm:"foreignKey:ApartmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"apartment"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Status string `gorm:"size:50;default:'active'"`

	Title string `gorm:"size:255;not null"`

	Pets           bool `gorm:"default:false"`
	Babies         bool `gorm:"default:false"`
	Smoking        bool `gorm:"default:false"`
	Internet       bool `gorm:"default:false"`
	WashingMachine bool `gorm:"default:false"`
	TV             bool `gorm:"default:false"`
	Conditioner    bool `gorm:"default:false"`
	Dishwasher     bool `gorm:"default:false"`
	Concierge      bool `gorm:"default:false"`

	Rent    float64 `gorm:"not null"`
	Deposit float64 `gorm:"not null"`

	RentalType string `gorm:"size:50"`
}
