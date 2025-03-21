package repository

import (
	"github.com/vasya/renting-app/internal/dto"
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
func (r *AuthPostgres) CreateUser(user dto.CreateUser) (int, error) {
	userGorm := models.User{
		Name:         user.Name,
		Surname:      user.Surname,
		Email:        user.Email,
		PasswordHash: user.Password,
		Birthdate:    user.Birthdate,
		Role:         user.Role,
	}
	result := r.db.Create(&userGorm)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(userGorm.ID), nil
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
