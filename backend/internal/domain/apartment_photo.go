package domain

type ApartmentPhoto struct {
	ID          uint      `gorm:"primaryKey"`
	ApartmentID uint      `gorm:"not null;index"`
	Apartment   Apartment `gorm:"foreignKey:ApartmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsCover     bool      `gorm:"size:50;default:false"`
	URL         string    `gorm:"not null"`
}
