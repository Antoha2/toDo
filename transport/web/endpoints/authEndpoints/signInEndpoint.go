package endpoints

import (
	"context"
	"log"

	authservice "github.com/antoha2/todo/service/authService"

	"github.com/go-kit/kit/endpoint"
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    []string
}

type SignInResponse struct {
	Token string `json:"token"`
}

func MakeSignInEndpoint(s authservice.Authorization) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(SignInRequest)

		token, err := s.GenerateToken(req.Username, req.Password)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return SignInResponse{Token: token}, nil
	}
}
