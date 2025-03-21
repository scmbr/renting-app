package repository

import (
	"errors"

	"github.com/vasya/renting-app/internal/dto"
	"github.com/vasya/renting-app/internal/models"
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
func (r *UsersPostgres) GetAllUsers() ([]dto.GetUser, error) {
	var users []models.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	var getUserDTOs []dto.GetUser
	for _, user := range users {
		getUserDTO := dto.GetUser{
			Id:        int(user.ID),
			Name:      user.Name,
			Surname:   user.Surname,
			Email:     user.Email,
			Birthdate: user.Birthdate,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			IsActive:  user.IsActive,
		}
		getUserDTOs = append(getUserDTOs, getUserDTO)
	}

	return getUserDTOs, nil
}

// GetUserById — получение пользователя по ID
func (r *UsersPostgres) GetUserById(id int) (*dto.GetUser, error) {
	var user models.User
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	getUserDTO := dto.GetUser{
		Id:        int(user.ID),
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		Birthdate: user.Birthdate,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		IsActive:  user.IsActive,
	}
	return &getUserDTO, nil
}

func (r *UsersPostgres) UpdateAvatar(userId int, avatarURL string) error {
	var user models.User
	result := r.db.First(&user, "id = ?", userId)
	if result.Error != nil {
		return result.Error
	}
	user.ProfilePicture = avatarURL
	result = r.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
