package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	renting_app "github.com/vasya/renting-app"
)

type AuthPostgres struct {
	db *sqlx.DB
}
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres{
	return &AuthPostgres{db:db}
}
func(r *AuthPostgres) CreateUser (user renting_app.User) (int,error){
	var id int
	query:= fmt.Sprintf("INSERT INTO %s (name,surname,email,password_hash,birthdate,role) values ($1,$2,$3,$4,$5,$6) RETURNING id", usersTable)
	row:=r.db.QueryRow(query,user.Name,user.Surname,user.Email,user.Password,user.Birthdate,user.Role)
	if err:=row.Scan(&id);err!=nil{
		return 0,err
	}
	return id,nil
}

func(r *AuthPostgres) GetUser(email,password string)(renting_app.User,error){
	var user renting_app.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err:=r.db.Get(&user,query,email,password)
	return user,err
}