package repository

import (
	"log"

	etodo "github.com/antoha2/todo"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	dbx *sqlx.DB
}

func NewAuthPostgres(dbx *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{dbx: dbx}
}

func (r *AuthPostgres) CreateUser(user *etodo.User) error {

	query := "INSERT INTO userlist (name, username, password) VALUES ($1, $2, $3) RETURNING id"
	row := r.dbx.QueryRow(query, user.Name, user.Username, user.Password)
	var id int
	if err := row.Scan(&id); err != nil {
		return err
	}
	user.Id = id
	log.Printf("создан пользователь id - %d, name - %s, username - %s", id, user.Name, user.Username)
	return nil
}
