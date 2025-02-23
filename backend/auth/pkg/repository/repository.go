package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vasya/auth"
)

type Authorization interface {
	CreateUser (user auth.User) (int,error)
}
type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}