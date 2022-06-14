package web

import (
	"net/http"

	authService "github.com/antoha2/todo/service/authService"
	taskService "github.com/antoha2/todo/service/taskService"
)

type Transport interface {
}

type webImpl struct {
	taskService taskService.TodolistService
	authService authService.AuthService
	server      *http.Server
}

func NewWeb(taskService taskService.TodolistService, authService authService.AuthService) *webImpl {
	return &webImpl{
		taskService: taskService,
		authService: authService,
	}
}
