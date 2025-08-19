package postgres

import (
	"fmt"

	"github.com/scmbr/renting-app/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Session{})
	db.AutoMigrate(&domain.Apartment{})
	db.AutoMigrate(&domain.Advert{})
	db.AutoMigrate(&domain.ApartmentPhoto{})
	db.AutoMigrate(&domain.Favorites{})
	db.AutoMigrate(&domain.Notification{})
	db.AutoMigrate(&domain.Review{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
