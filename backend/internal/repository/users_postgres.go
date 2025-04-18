package repository

import (
	"errors"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/models"
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
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var user models.User
	result := tx.First(&user, "id = ?", id)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("user not found")
	}
	tx.Commit()
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
func (r *UsersPostgres) DeleteUserById(id int) (*dto.GetUser, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var user models.User
	result := tx.First(&user, "id = ?", id)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("user not found")
	}
	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
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
	tx.Commit()
	return &getUserDTO, nil
}
func (r *UsersPostgres) UpdateUserById(input *dto.GetUser) (*dto.GetUser, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var user models.User
	result := tx.First(&user, "id = ?", input.Id)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("user not found")
	}
	user.Name = input.Name
	user.Surname = input.Surname
	user.Email = input.Email
	user.Birthdate = input.Birthdate
	user.Role = input.Role
	user.IsActive = input.IsActive

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
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
	tx.Commit()
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
