package endpoints

import (
	authservice "github.com/antoha2/todo/service/authService"
	service "github.com/antoha2/todo/service/taskService"

	authEndpoints "github.com/antoha2/todo/transport/web/endpoints/authEndpoints"
	taskEndpoints "github.com/antoha2/todo/transport/web/endpoints/taskEndpoints"
	"github.com/go-kit/kit/endpoint"
)

type TaskEndpoints struct {
	Create endpoint.Endpoint
	Read   endpoint.Endpoint
	Update endpoint.Endpoint
	Delete endpoint.Endpoint
}

func MakeTaskEndpoints(s service.TodolistService) *TaskEndpoints {
	return &TaskEndpoints{
		Create: taskEndpoints.MakeCreateEndpoint(s),
		Read:   taskEndpoints.MakeReadEndpoint(s),
		Update: taskEndpoints.MakeUpdateEndpoint(s),
		Delete: taskEndpoints.MakeDeleteEndpoint(s),
	}
}

type AuthEndpoints struct {
	SignIn      endpoint.Endpoint
	SignUpAdmin endpoint.Endpoint
	SignUpUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
}

func MakeAuthEndpoints(s authservice.Authorization) *AuthEndpoints {
	return &AuthEndpoints{
		SignIn:      authEndpoints.MakeSignInEndpoint(s),
		SignUpAdmin: authEndpoints.MakeSignUpAdminEndpoint(s),
		SignUpUser:  authEndpoints.MakeSignUpUserEndpoint(s),
		DeleteUser:  authEndpoints.MakeDeleteUserEndpoint(s),
		UpdateUser:  authEndpoints.MakeUpdateUserEndpoint(s),
	}

}
