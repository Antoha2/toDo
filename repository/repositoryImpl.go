package repository

import "fmt"

type repositoryImpl struct {
	rep []RepTask
}

func New() *repositoryImpl {
	r := make([]RepTask, 0)

	return &repositoryImpl{
		rep: r,
	}
}

func (r *repositoryImpl) Create(task *RepTask) error {

	r.rep = append(r.rep, *task)
	fmt.Println(r)
	return nil
}
