package repository

import (
	"context"
	"errors"

	"github.com/scmbr/renting-app/internal/dto"
	"github.com/scmbr/renting-app/internal/models"
	"gorm.io/gorm"
)

// UsersRepo — структура для работы с пользователями в PostgreSQL через GORM
type UsersRepo struct {
	db *gorm.DB
}

// NewUsersRepo — конструктор для UsersRepo
func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

// GetAllUsers — получение всех пользователей
func (r *UsersRepo) GetAllUsers() ([]dto.GetUser, error) {
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
func (r *UsersRepo) GetUserById(id int) (*dto.GetUser, error) {
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
		Id:               int(user.ID),
		Name:             user.Name,
		Surname:          user.Surname,
		Email:            user.Email,
		Birthdate:        user.Birthdate,
		Role:             user.Role,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
		VerificationCode: user.VerificationCode,
		Verified:         user.Verified,
		IsActive:         user.IsActive,
	}
	return &getUserDTO, nil
}
func (r *UsersRepo) DeleteUserById(id int) (*dto.GetUser, error) {
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
func (r *UsersRepo) UpdateUserById(input *dto.GetUser) (*dto.GetUser, error) {
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
func (r *UsersRepo) UpdateAvatar(userId int, avatarURL string) error {
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

func (r *UsersRepo) CreateUser(user dto.CreateUser, code string) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	userGorm := models.User{
		Name:             user.Name,
		Surname:          user.Surname,
		Email:            user.Email,
		PasswordHash:     user.Password,
		Birthdate:        user.Birthdate,
		VerificationCode: code,
	}
	result := tx.Create(&userGorm)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}
func (r *UsersRepo) GetByCredentials(ctx context.Context, email, passwordHash string) (*dto.GetUser, error) {

	var user models.User
	result := r.db.WithContext(ctx).Where("email = ? AND password_hash = ?", email, passwordHash).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
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

// GetUser — получает пользователя по email и паролю
func (r *UsersRepo) GetUser(email, password string) (models.User, error) {
	var user models.User
	result := r.db.Where("email = ? AND password_hash = ?", email, password).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
func (r *UsersRepo) Verify(ctx context.Context, code string) (dto.GetUser, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Where("verification_code = ?", code).
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.GetUser{}, errors.New("user with the given verification code not found")
		}
		return dto.GetUser{}, err
	}

	user.Verified = true
	user.VerificationCode = ""

	err = r.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return dto.GetUser{}, err
	}

	return dto.GetUser{
		Id:               int(user.ID),
		Name:             user.Name,
		Surname:          user.Surname,
		Email:            user.Email,
		Birthdate:        user.Birthdate,
		Role:             user.Role,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
		VerificationCode: user.VerificationCode,
		Verified:         user.Verified,
		IsActive:         user.IsActive,
	}, nil
}
