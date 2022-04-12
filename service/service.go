package service

import _ "github.com/antoha2/todo/repository"

type Service interface {
	Create(task *SerTask) error
	Read(task *SerFilter) []*SerTask
	Delete(task *SerTask) error
	Update(task *SerTask) error
	LenRep() int
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
