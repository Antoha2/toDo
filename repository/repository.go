package repository

import (
	etodo "github.com/antoha2/todo"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	TodolistRep
}

type Authorization interface {
	CreateUser(user *etodo.User) error
}

type TodolistRep interface {
	Create(*RepTask) error
	Read(*RepFilter) []RepTask
	Update(*RepTask) error
	Delete(*RepTask) error //Delete(*RepFilter) error
}

type repositoryImplDB struct {
	rep *sqlx.DB
}

func NewDB(dbx *sqlx.DB) *repositoryImplDB {

	return &repositoryImplDB{
		rep: dbx,
	}
}

func NewRepository(dbx *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthPostgres(dbx),
		TodolistRep:   NewDB(dbx),
	}
}
