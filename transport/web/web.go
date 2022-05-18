package web

import (
	"context"
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

func NewContext(ctx context.Context, userId int) context.Context {
	return context.WithValue(ctx, ContextUserId, userId)
}

func FromContext(ctx context.Context) (int, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// type assertion returns ok=false for nil.
	userId, ok := ctx.Value(ContextUserId).(int)
	return userId, ok
}

type Task struct {
	Id     int    `json:"task_id"`
	UserId int    `json:"user_id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}
