package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	renting_app "github.com/vasya/renting-app"
)

type UsersPostgres struct {
	db *sqlx.DB
}
func NewUsersPostgres(db *sqlx.DB) *UsersPostgres{
	return &UsersPostgres{db:db}
}


func (r *UsersPostgres) GetAllUsers() ([]renting_app.GetUser, error) {
    var users []renting_app.GetUser
    query := fmt.Sprintf("SELECT id, name,surname,email,birthdate,role,created_at,updated_at,is_active FROM %s", usersTable)
    
    err := r.db.Select(&users, query)
    if err != nil {
        return nil, err
    }
    
    return users, nil
}

func (r *UsersPostgres) GetUserById(id int) (*renting_app.GetUser, error) {
    var user renting_app.GetUser
    query := fmt.Sprintf("SELECT id, name,surname,email,birthdate,role,created_at,updated_at,is_active FROM %s WHERE id=$1", usersTable)
    
    err := r.db.Get(&user, query, id)
    if err != nil {
        return nil, err
    }
    
    return &user, nil
}


