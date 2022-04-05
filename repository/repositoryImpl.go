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

func (r *repositoryImpl) Read(readTask *RepFilter) *RepTask {

	//fmt.Println(r.rep[i-1])
	//tsd := r.rep[i-1]

	var tsd RepTask
	for i, _ := range r.rep {
		if r.rep[i].Id == readTask.Id {
			tsd = r.rep[readTask.Id-1]

		}
	}
	fmt.Println(tsd)
	return &tsd

}
