package endpoints

import (
	"context"

	etodo "github.com/antoha2/todo"
	authservice "github.com/antoha2/todo/service/authService"
	"github.com/go-kit/kit/endpoint"
)

type SignUpUserRequest struct {
	FirstName string `json:"firstname" gorm:"column:firstname"`
	LastName  string `json:"lastname" gorm:"column:lastname"`
	Username  string `json:"username" gorm:"column:username"`
	Password  string `json:"password"`
}

type SignUpUserResponse struct {
}

func MakeSignUpUserEndpoint(s authservice.Authorization) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(SignUpUserRequest)

		inputUser := new(etodo.User)
		inputRoles := new(etodo.UsersRoles)

		inputUser.FirstName = req.FirstName
		inputUser.LastName = req.LastName
		inputUser.Password = req.Password
		inputUser.Username = req.Username

		inputRoles.Roles = append(inputRoles.Roles, roleUser)

		if err := s.CreateUser(inputUser, inputRoles); err != nil {

			return nil, err
		}

		return SignUpUserResponse{}, nil
	}
}
