package service

import (
	"fmt"
	"time"

	taskRepository "github.com/antoha2/todo/repository"
)

//Create
func (s *serviceImpl) Create(task *SerTask) error {

	repTask := new(taskRepository.RepTask)
	repTask.Text = task.Text
	repTask.CreateAt = time.Now()
	repTask.IsDone = false

	err := s.repository22.Create(repTask)
	if err != nil {
		fmt.Println(err)
		return err
	}
	task.Id = repTask.Id
	task.IsDone = repTask.IsDone

	return nil
}

//Read
func (s *serviceImpl) Read(task *SerFilter) []*SerTask {

	readFilter := new(taskRepository.RepFilter)
	readFilter.Ids = task.Ids
	tasks := s.repository22.Read(readFilter)
	sliceTask := make([]*SerTask, len(tasks))

	for index, task := range tasks {
		t := &SerTask{
			Id:     task.Id,
			Text:   task.Text,
			IsDone: task.IsDone,
		}
		sliceTask[index] = t
	}
	return sliceTask
}

//Delete
func (s *serviceImpl) Delete(task *SerTask) error {

	//delFilter := new(repository.RepFilter)
	delFilter := new(taskRepository.RepTask)
	delFilter.Id = task.Id
	err := s.repository22.Delete(delFilter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//Update
func (s *serviceImpl) Update(task *SerTask) error {

	upFilter := new(taskRepository.RepTask)
	upFilter.Id = task.Id
	upFilter.IsDone = task.IsDone
	upFilter.Text = task.Text

	err := s.repository22.Update(upFilter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//task.IsDone = repTask.IsDone
	//	task.Text = repTask.Text

	return nil
}
