package domain

type Room struct {
	ID          uint      `gorm:"primaryKey"`
	ApartmentID uint      `gorm:"not null;index"`
	Apartment   Apartment `gorm:"foreignKey:ApartmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      *uint
	User        *User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Area        float64 `gorm:"not null"`
	Status      string  `gorm:"size:50;default:'vacant'"`
}
