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

func (r *repositoryImpl) Read(i int) RepTask {

	//r.rep = append(r.rep, *task)
	//fmt.Println(r.rep[i-1])
	//tsd:=r.rep[i-1]
	fmt.Println(r.rep[i-1])
	return r.rep[i-1]
}
