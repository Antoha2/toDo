package endpoints

import (
	"context"

	etodo "github.com/antoha2/todo"
	authservice "github.com/antoha2/todo/service/authService"
	"github.com/go-kit/kit/endpoint"
)

type SignUpAdminRequest struct {
	FirstName string `json:"firstname" gorm:"column:firstname"`
	LastName  string `json:"lastname" gorm:"column:lastname"`
	Username  string `json:"username" gorm:"column:username"`
	Password  string `json:"password"`
}

type SignUpAdminResponse struct {
}

const (
	roleAdmin = "admin"
	roleUser  = "user"
	roleDev   = "dev"
)

func MakeSignUpAdminEndpoint(s authservice.Authorization) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(SignUpAdminRequest)

		inputUser := new(etodo.User)
		inputRoles := new(etodo.UsersRoles)

		inputUser.FirstName = req.FirstName
		inputUser.LastName = req.LastName
		inputUser.Password = req.Password
		inputUser.Username = req.Username

		inputRoles.Roles = append(inputRoles.Roles, roleAdmin)
		inputRoles.Roles = append(inputRoles.Roles, roleDev)

		if err := s.CreateUser(inputUser, inputRoles); err != nil {

			return nil, err
		}

		return SignUpAdminResponse{}, nil
	}
}
