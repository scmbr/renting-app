package repository

import (
	"fmt"

	"github.com/scmbr/renting-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config — структура для конфигурации подключения к PostgreSQL
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB — создает подключение к PostgreSQL через GORM
func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Session{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
