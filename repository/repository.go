package repository

import "time"

type Repository interface {
	Create(*RepTask) error
	Read(*RepFilter) *RepTask
	Delete(*RepFilter) error
	/* Read() []RepTask
	Update(RepTask) RepTask
	*/

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
