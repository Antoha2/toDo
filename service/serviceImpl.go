package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/antoha2/todo/repository"
)

type serviceImpl struct {
	repository22 repository.Repository
	counter      func() int
}

func New(rep repository.Repository) *serviceImpl {
	counter := newCounter()
	return &serviceImpl{
		repository22: rep,
		counter:      counter,
	}
}

//Create
func (s *serviceImpl) Create(task *SerTask) error {

	if s.repository22.LenRep() > 3 {
		return errors.New("нельзя хранить больше трех")
	}
	newId := s.counter()
	repTask := new(repository.RepTask)

	repTask.Text = task.Text
	repTask.Id = newId
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

//счетчик уникальных Id
func newCounter() func() int {
	var count int
	couner := func() int {
		count++
		return count
	}
	return couner
}

//Read
func (s *serviceImpl) Read(task *SerFilter) *SerFilter {

	readFilter := new(repository.RepFilter)

	//fmt.Println(readFilter)
	readFilter.Ids = task.Ids
	//fmt.Println(readFilter)
	//repTask :
	_ = s.repository22.Read(readFilter)

	//fmt.Println(repTask.Tasks)
	//task.Tasks = repTask.Tasks
	/* task.IsDone = repTask.IsDone
	task.Text = repTask.Text */

	return task
}

//Delete
func (s *serviceImpl) Delete(task *SerTask) error {

	delFilter := new(repository.RepFilter)
	delFilter.Id = task.Id
	err := s.repository22.Delete(delFilter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *serviceImpl) LenRep() int {

	return 0
}

//Update
func (s *serviceImpl) Update(task *SerTask) error {

	upFilter := new(repository.RepTask)
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
