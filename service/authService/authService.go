package authservice

import (
	"time"

	etodo "github.com/antoha2/todo"

	authRepository "github.com/antoha2/todo/service/authService/authRepository"
)

const (
	salt       = "aW1;"
	signingKey = "Bgt5"
	tokenTTL   = 12 * time.Hour
)

type Authorization interface {
	CreateUser(user *etodo.User) error
	UpdateUser(user *etodo.User) error
	DeleteUser(user *etodo.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type AuthService struct {
	authRep authRepository.Authorization
}

func NewAuthService(authRep authRepository.Authorization) *AuthService {
	return &AuthService{
		authRep: authRep,
	}
}
