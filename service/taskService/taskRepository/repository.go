package repository

import (
	"gorm.io/gorm"
)

type taskRepository struct {
	TodolistRep
}

type TodolistRep interface {
	Create(*RepTask) error
	Read(*RepFilter) []RepTask
	Update(*RepTask) error
	Delete(*RepTask) error //Delete(*RepFilter) error
}

type RepTask struct {
	TaskId int `gorm:"primaryKey;"` //  index:unique"`
	UserId int
	Text   string
	IsDone bool `gorm:"column:isdone"`
	//CreateAt time.Time
}

func (RepTask) TableName() string {
	return "todolist"
}

type RepFilter struct {
	TaskId    int
	UserId    int
	UserRoles []string
	Ids       []int
	Text      string
	IsDone    bool
	//Tasks  []RepTask
}

type repositoryImplDB struct {
	rep *gorm.DB
}

func NewDB(dbx *gorm.DB) *repositoryImplDB {

	return &repositoryImplDB{
		rep: dbx,
	}
}

func NewTaskRepository(dbx *gorm.DB) *taskRepository {

	return &taskRepository{
		TodolistRep: NewDB(dbx),
	}
}
