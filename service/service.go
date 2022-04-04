package service

import _ "github.com/staszigzag/etodo/repository"

type Service interface {
	Create(task *SerTask) error
}

type SerTask struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}
