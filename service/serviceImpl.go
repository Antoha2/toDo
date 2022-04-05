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

func (s *serviceImpl) Create(task *SerTask) error {

	newId := s.counter()
	if newId > 3 {
		return errors.New("нельзя хранить больше трех")
	}
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

func newCounter() func() int {
	var count int
	couner := func() int {
		count++
		return count
	}
	return couner
}
