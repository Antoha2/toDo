package repository

import "time"

type Repository interface {
	Create(*RepTask) error
	Read(*RepFilter) *RepTask
	Delete(*RepFilter) error
	LenRep() int
	Update(*RepTask) error //*RepTask
}

type RepTask struct {
	Id       int
	Text     string
	IsDone   bool
	CreateAt time.Time
}

type RepFilter struct {
	Id int
}
