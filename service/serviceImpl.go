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

func (s *serviceImpl) Read(task SerTask) SerTask {

	//s.rep
	i := task.Id
	//repTask := new(repository.RepTask)

	//if i <1 || i>len(s.){

	//i, err := strconv.Atoi(task.Id)
	/*
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		} */
	//fmt.Println("serv1")

	repTask := s.repository22.Read(i)

	task.IsDone = repTask.IsDone
	task.Text = repTask.Text

	/*  err := s.repository22.Read(i)
	if err != nil {
		fmt.Println(err)
		return err
	} */
	//fmt.Println("serv2")

	return task
}
