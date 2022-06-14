package endpoints

import (
	"context"
	"errors"
	"log"

	etodo "github.com/antoha2/todo"
	authservice "github.com/antoha2/todo/service/authService"
	"github.com/go-kit/kit/endpoint"
)

type UpdateUserRequest struct {
	FirstName string `json:"firstname" gorm:"column:firstname"`
	LastName  string `json:"lastname" gorm:"column:lastname"`
	Password  string `json:"password"`
}

type UpdateUserResponse struct {
	UserId int `json:"user_id"`
}

func MakeUpdateUserEndpoint(s authservice.Authorization) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(UpdateUserRequest)
		user := new(etodo.User)

		var userId int
		var ok bool

		if userId, ok = ctx.Value(etodo.USER_ID).(int); !ok {
			newErr := "UserId не найден"
			log.Println(newErr)
			return nil, errors.New(newErr)
		}
		user.UserId = userId
		user.FirstName = req.FirstName
		user.LastName = req.LastName
		user.Password = req.Password

		if err := s.UpdateUser(user); err != nil {
			return nil, err
		}
		return UpdateUserResponse{UserId: userId}, nil
	}
}
