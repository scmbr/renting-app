package repository

import (
	"github.com/vasya/renting-app/internal/dto"
	"github.com/vasya/renting-app/internal/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser (user models.User) (int,error)
	GetUser(email,password string) (models.User,error)
}
type Users interface{
	GetAllUsers() ([]dto.GetUser,error)
	GetUserById(id int) (*dto.GetUser,error)
}
type Repository struct {
	Authorization
	Users
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Users: NewUsersPostgres(db),
	}
}