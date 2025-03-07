package repository

import (
	"github.com/vasya/renting-app/internal/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser — создает нового пользователя в базе данных
func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(user.Id), nil
}

// GetUser — получает пользователя по email и паролю
func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	result := r.db.Where("email = ? AND password_hash = ?", email, password).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
