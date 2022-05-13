package authservice

import (
	"crypto/sha1"
	"fmt"
	"log"

	etodo "github.com/antoha2/todo"

	taskRepository "github.com/antoha2/todo/repository"
)

const salt = "aW1;" //"HjjKcxzo2j1J11p"

type Authorization interface {
	CreateUser(user *etodo.User) error
}

type AuthService struct {
	authRep taskRepository.Authorization
}

func NewAuthService(authRep taskRepository.Authorization) *AuthService {
	return &AuthService{
		authRep: authRep,
	}
}

func (s *AuthService) CreateUser(user *etodo.User) error {

	user.Password = s.generatePasswordHash(user.Password)
	err := s.authRep.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Println(user)
	return nil
}

//шифрование пароля
func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
