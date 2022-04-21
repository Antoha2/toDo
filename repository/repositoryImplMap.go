package repository

import (
	"errors"
	"fmt"
)

var errNotFinedIdMap = errors.New("id not fined")

type repositoryImplMap struct {
	rep map[int]RepTask
}

func NewMap() *repositoryImplMap {
	r := make(map[int]RepTask)

	return &repositoryImplMap{
		rep: r,
	}
}

//Create
func (r *repositoryImplMap) Create(task *RepTask) error {

	r.rep[task.Id] = *task
	//fmt.Println(r)
	return nil
}

//подсчет кол-ва элементов
func (r *repositoryImplMap) LenRep() int {

	count := len(r.rep)
	return count
}

//Read
func (r *repositoryImplMap) Read(readFilter *RepFilter) []RepTask {

	sliceTask := make([]RepTask, 0)

	if readFilter.Ids == nil || len(readFilter.Ids) == 0 {

		for _, task := range r.rep {
			sliceTask = append(sliceTask, task)
		}
		return sliceTask
	}

	for _, id := range readFilter.Ids {
		for i, task := range r.rep {
			if id == task.Id {
				sliceTask = append(sliceTask, r.rep[i])
			}
		}
	}

	for _, t := range sliceTask {
		fmt.Println(t)
	}

	return sliceTask
}

//Delete
func (r *repositoryImplMap) Delete(delFilter *RepFilter) error {

	if _, ok := r.rep[delFilter.Id]; ok {
		delete(r.rep, delFilter.Id)
		fmt.Println(r)
		return nil
	}
	return errNotFinedIdMap
}

//Update
func (r *repositoryImplMap) Update(upTask *RepTask) error {

	if _, ok := r.rep[upTask.Id]; ok {
		r.rep[upTask.Id] = *upTask
		fmt.Println(r)
		return nil
	}
	return errNotFinedIdMap
}
