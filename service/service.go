package service

import (
	//etodo "github.com/antoha2/todo"

	taskRepository "github.com/antoha2/todo/repository"
	authservice "github.com/antoha2/todo/service/authService"
	// "github.com/antoha2/todo/service/authservice"
)

type Service struct {
	authservice.Authorization
	TodolistServ
}

type TodolistServ interface {
	Create(task *SerTask) error
	Read(task *SerFilter) []*SerTask
	Delete(task *SerTask) error
	Update(task *SerTask) error
}

type serviceImpl struct {
	repository22 taskRepository.TodolistRep
}

func NewTaskService(rep taskRepository.TodolistRep) *serviceImpl {
	return &serviceImpl{
		repository22: rep,
	}
}

type SerTask struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}

type SerFilter struct {
	Id     int    `json:"id"`
	Ids    []int  `json:"ids"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
	//Tasks  []SerTask `json:"tasks"`
}
