package repository

import (
	"errors"
	"fmt"
)

type repositoryImplMap struct {
	rep map[int]RepTask //////////////////////////////////////////////////////////////////
}

func NewMap() *repositoryImplMap {
	r := make(map[int]RepTask)

	return &repositoryImplMap{
		rep: r,
	}
}

func (r *repositoryImplMap) Create(task *RepTask) error {

	//t:=task.Id
	r.rep[task.Id] = *task
	//r.rep[t] = append(r.rep, *task)
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

	delete(r.rep, delFilter.Id)
	fmt.Println(r)
	return nil

	/* if delFilter.Ids == nil || len(delFilter.Ids) == 0 {
		for i, _ := range r.rep {
			delete(r.rep, i)

		}
		fmt.Println(r)
		return nil
	}  */
	//for _, id := range delFilter.Ids {

	//fmt.Println(len(delFilter.Ids))

}

//Update
func (r *repositoryImplMap) Update(upTask *RepTask) error {

	sliceTask := make([]RepTask, 0)
	for _, task := range r.rep {
		sliceTask = append(sliceTask, task)
	}

	isUpdate := false

	for index, _ := range sliceTask {

		if sliceTask[index].Id == upTask.Id {

			sliceTask[index].Text = upTask.Text
			sliceTask[index].IsDone = upTask.IsDone
			fmt.Println(sliceTask[index])
			isUpdate = true
		}
	}

	for index, task := range sliceTask {

		r.rep[index+1] = task
	}

	if !isUpdate {
		return errors.New("id not fined")
	}
	return nil
}
