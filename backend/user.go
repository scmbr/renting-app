package renting_app

import "time"

type User struct {
	Id        int       `json:"-" db:"id"`
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json:"surename" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Birthdate time.Time `json:"birthdate" binding:"required"`
	Role      int       `json:"role" binding:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	IsActive  bool `json:"-"`
}

type GetUser struct{
	Id        int       `json:"-" db:"id"`
	Name      string    `json:"name" binding:"required" db:"name"` 
	Surname   string    `json:"surename" binding:"required" db:"surname"`
	Email     string    `json:"email" binding:"required" db:"email"`
	Birthdate time.Time `json:"birthdate" binding:"required" db:"birthdate"`
	Role      int       `json:"role" binding:"required" db:"role"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
	IsActive  bool `json:"-" db:"is_active"`
}

// CREATE TABLE users
// (
//     id SERIAL NOT NULL UNIQUE,
//     name VARCHAR(255) NOT NULL,
//     surname VARCHAR(255) NOT NULL,
//     email VARCHAR(255) NOT NULL,
//     password_hash VARCHAR(255) NOT NULL,
//     birthdate DATE NOT NULL,
//     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
//     role INT NOT NULL,
//     is_active BOOLEAN DEFAULT TRUE
// );