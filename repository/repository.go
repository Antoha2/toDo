package repository

import "time"

type Repository interface {
	Create(*RepTask) error
	/* Read() []RepTask
	Update(RepTask) RepTask
	Delete(RepTask) */

}

type RepTask struct {
	Id       int
	Text     string
	IsDone   bool
	CreateAt time.Time
}
