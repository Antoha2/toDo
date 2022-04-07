package repository

import (
	"errors"
	"fmt"
)

type repositoryImplMap struct {
	rep mapRepTask //////////////////////////////////////////////////////////////////
}

func NewMap() *repositoryImplMap {
	r := make([]RepTask, 0)

	return &repositoryImplMap{
		rep: r,
	}
}

func (r *repositoryImplMap) Create(task *RepTask) error {

	r.rep = append(r.rep, *task)
	fmt.Println(r)
	return nil
}

//подсчет кол-ва элементов
func (r *repositoryImplMap) LenRep() int {

	count := len(r.rep)
	return count
}

//Read
func (r *repositoryImplMap) Read(readFilter *RepFilter) []RepTask {

	if readFilter.Ids == nil || len(readFilter.Ids) == 0 {
		return r.rep
	}

	sliceTask := make([]RepTask, 0)

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
func (r *repositoryImplMap) Delete(delTask *RepFilter) error {

	for i, v := range r.rep {
		if v.Id == delTask.Id {
			copy(r.rep[i:], r.rep[i+1:])
			//r.rep[len(r.rep)-1] = nil // обнуляем "хвост"
			r.rep = r.rep[:len(r.rep)-1]
		}
	}
	fmt.Println(r)
	return nil
}

//Update
func (r *repositoryImplMap) Update(upTask *RepTask) error {

	isUpdate := false

	for index, _ := range r.rep {

		if r.rep[index].Id == upTask.Id {

			r.rep[index].Text = upTask.Text
			r.rep[index].IsDone = upTask.IsDone
			fmt.Println(r.rep[index])
			isUpdate = true
		}
	}
	if !isUpdate {
		return errors.New("id not fined")
	}
	return nil
}