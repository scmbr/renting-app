package models

import (
	"time"
)

type Favorites struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`          
    AdvertID uint      `gorm:"not null"`         
    CreatedAt time.Time `gorm:"autoCreateTime"`   
    User    User    `gorm:"foreignKey:UserID"`
    Advert Advert `gorm:"foreignKey:AdvertID"`
}
