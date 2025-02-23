package auth

import "time"

type User struct {
	Id        int       `json: "-"`
	Name      string    `json:"name" binding:"required"`
	Surname   string    `json: "surename" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Birthdate time.Time `json:"birthdate" binding:"required"`
	Role      int       `json:"role" binding:"required"`
	CreatedAt time.Time `json: "-"`
	UpdatedAt time.Time `json: "-"`
	IsActive  time.Time `json: "-"`
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