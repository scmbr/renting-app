package domain

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	AuthorID  uint      `gorm:"not null"`
	Author    User      `gorm:"foreignKey:AuthorID"`
	TargetID  uint      `gorm:"not null"`
	Target    User      `gorm:"foreignKey:TargetID"`
	Rating    int       `gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Comment   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
