package endpoints

import (
	"context"
	"errors"
	"fmt"
	"log"

	etodo "github.com/antoha2/todo"
	service "github.com/antoha2/todo/service/taskService"
	"github.com/go-kit/kit/endpoint"
)

type CreateRequest struct {
	Text string `json:"text"`
}

type CreateResponse struct {
	//TaskId int `json:"task_id"`
	str string `json:"text"`
}

/*  type Task struct {
	TaskId int    `json:"id"`
	UserId int    `json:"user_id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}
*/
func MakeCreateEndpoint(s service.TodolistService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(CreateRequest)
		serTask := new(service.SerTask)
		serTask.Text = req.Text

		var ok bool
		if serTask.UserId, ok = ctx.Value(etodo.USER_ID).(int); !ok {
			newErr := "UserId не найден"
			log.Println(newErr)
			return nil, errors.New(newErr)
		}
		if err := s.Create(ctx, serTask); err != nil {
			return nil, err
		}

		task := &etodo.Task{
			Id:     serTask.Id,
			UserId: serTask.UserId,
			Text:   serTask.Text,
			IsDone: serTask.IsDone,
		}
		//respTasks[index] = t
		//log.Println(task)
		t := fmt.Sprintf("создана запись - %v", task)
		//log.Println(t)
		return CreateResponse{str: t}, nil
	}
}
