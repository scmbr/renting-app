package repository

import (
	renting_app "github.com/vasya/renting-app"
	"gorm.io/gorm"
)

// UsersPostgres — структура для работы с пользователями в PostgreSQL через GORM
type UsersPostgres struct {
	db *gorm.DB
}

// NewUsersPostgres — конструктор для UsersPostgres
func NewUsersPostgres(db *gorm.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

// GetAllUsers — получение всех пользователей
func (r *UsersPostgres) GetAllUsers() ([]renting_app.GetUser, error) {
	var users []renting_app.GetUser
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// GetUserById — получение пользователя по ID
func (r *UsersPostgres) GetUserById(id int) (*renting_app.GetUser, error) {
	var user renting_app.GetUser
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

