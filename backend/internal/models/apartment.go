package models

import "time"

type Apartment struct {
	ID               uint      `gorm:"primaryKey"`
	UserID           uint      `gorm:"not null;index"` // внешний ключ
	User             User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	City             string    `gorm:"size:100;not null"`
	CitySlug         string    `gorm:"size:50;default:'kazan'"`
	Street           string    `gorm:"size:100;not null"`
	District         string    `gorm:"size:100"`
	House            string    `gorm:"size:20"`
	Building         string    `gorm:"size:20"`
	Floor            int       `gorm:"not null"`
	ApartmentNumber  string    `gorm:"size:20"`
	Longitude        float64   `gorm:"not null"`
	Latitude         float64   `gorm:"not null"`
	Rooms            int       `gorm:"not null"`
	Elevator         bool      `gorm:"default:false"`
	GarbageChute     bool      `gorm:"default:false"`
	BathroomType     string    `gorm:"size:50"`
	Concierge        bool      `gorm:"default:false"`
	ConstructionYear int       `gorm:"not null"`
	ConstructionType string    `gorm:"size:100"`
	Remont           string    `gorm:"size:100"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
	Rating           float32   `gorm:"default:0"`
	Status           string    `gorm:"size:50;default:'active'"`
}
