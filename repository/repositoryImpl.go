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

func (r *repositoryImpl) Read(readTask *RepReadTask) *RepTask {

	//fmt.Println(r.rep[i-1])
	//tsd := r.rep[i-1]

	var tsd RepTask
	for id, _ := range r.rep {
		if r.rep[id].Id == readTask.Id {
			tsd = r.rep[readTask.Id-1]

		}
	}
	fmt.Println(tsd)
	return &tsd

}
