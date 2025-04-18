package repository

import (
	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/models"
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
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	userGorm := models.User{
		Name:         user.Name,
		Surname:      user.Surname,
		Email:        user.Email,
		PasswordHash: user.Password,
		Birthdate:    user.Birthdate,
	}
	result := tx.Create(&userGorm)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}
	tx.Commit()
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
