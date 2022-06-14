package authservice

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	etodo "github.com/antoha2/todo"
	"github.com/dgrijalva/jwt-go"
)

func (s *AuthService) CreateUser(user *etodo.User, userRoles *etodo.UsersRoles) error {

	user.Password = s.generatePasswordHash(user.Password)
	err := s.authRep.CreateUser(user, userRoles)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *AuthService) DeleteUser(userId int) error {
	err := s.authRep.DeleteUser(userId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *AuthService) UpdateUser(user *etodo.User) error {

	if user.Password != "" {
		user.Password = s.generatePasswordHash(user.Password)
	}

	err := s.authRep.UpdateUser(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//шифрование пароля
func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {

	user, err := s.authRep.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.UserId,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accesToken string) (int, error) {

	token, err := jwt.ParseWithClaims(accesToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неверный метод подписи")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims not of type *tokenClaims")
	}
	return claims.UserId, nil
}

//получение роли из бд для middleware
func (s *AuthService) GetRoles(userId int) []string {
	return s.authRep.GetRoles(userId)
}
