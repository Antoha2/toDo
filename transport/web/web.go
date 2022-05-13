package web

import (
	"net/http"

	taskService "github.com/antoha2/todo/service"
	authService "github.com/antoha2/todo/service/authService"
)

type Transport interface {
}
type Task struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}

type webImpl struct {
	taskService taskService.TodolistServ
	server      *http.Server
}

type authWebImpl struct {
	authService authService.Authorization
	server      *http.Server
}

type WebAuthUser struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
