package service

import _ "github.com/antoha2/todo/repository"

type Service interface {
	Create(task *SerTask) error
	Read(task *SerTask) *SerTask
	Delete(task *SerTask) error
	Update(task *SerTask) error //*SerTask
	LenRep() int
}

type SerTask struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}

type SerFilter struct {
	Id int `json:"id"`
}
