package authrepository

import (
	etodo "github.com/antoha2/todo"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *etodo.User, userRoles *etodo.UsersRoles) error
	UpdateUser(user *etodo.User) error
	DeleteUser(userId int) error
	GetUser(username, password string) (*etodo.User, error)
	GetRoles(id int) []string
}

type AuthPostgres struct {
	dbx *gorm.DB
}

func NewAuthPostgres(dbx *gorm.DB) *AuthPostgres {
	return &AuthPostgres{dbx: dbx}
}

/* type UserlistToRoles struct {
	user_id int
	role_id int
} */
