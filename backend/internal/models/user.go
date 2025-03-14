package models

import (
	"time"
)

type User struct {
	Id        int       
	Name      string    
	Surname   string    
	Email     string    
	PasswordHash  string   
	Birthdate time.Time 
	Role      int       
	ProfilePicture string
	CreatedAt time.Time 
	UpdatedAt time.Time 
	IsActive  bool 
}


