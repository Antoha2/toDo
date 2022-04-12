package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/antoha2/todo/repository"
	"github.com/antoha2/todo/service"
)

func InitTest() service.Service {
	rep := repository.NewMap()
	ser := service.New(rep)
	return ser
}

func InitCreate(s service.Service) {
	defCreate := []struct {
		create *service.SerTask
	}{
		{&service.SerTask{1, "111111", false}},
		{&service.SerTask{2, "222222", false}},
		{&service.SerTask{3, "333333", false}},
	}
	for _, task := range defCreate {
		err := s.Create(task.create)
		if err != nil {
			fmt.Printf("Create() error = %v", err)
		}
	}
}

//test Create
/* func TestCreateTask(t *testing.T) {
	tests := []struct {
		name    string
		input   *service.SerTask
		want    *service.SerTask
		wantErr bool
	}{
		{"создаю первую задачу", &service.SerTask{0, "111111", false}, &service.SerTask{1, "111111", false}, false},
		{"создаю вторую задачу", &service.SerTask{0, "222222", true}, &service.SerTask{2, "222222", false}, false},
		{"создаю третью задачу", &service.SerTask{0, "333333", true}, &service.SerTask{3, "333333", false}, false},
		{"создаю четвертую задачу", &service.SerTask{0, "444444", true}, &service.SerTask{0, "444444", true}, true},
	}

	s := InitTest()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := s.Create(tc.input)

			if (err != nil) != tc.wantErr { // если ошибка не нил , и не ждем ошибку , то ...
				t.Fatalf("Create() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(*tc.want, *tc.input) { //если нет ошибки , то сравниваем значения
				t.Fatalf("expected: %v, got: %v", *tc.want, *tc.input)
			}
		})
	}
} */

//test Read
func TestReadTask(t *testing.T) {

	defCreate := []struct {
		create *service.SerTask
	}{
		{&service.SerTask{1, "111111", false}},
		{&service.SerTask{2, "222222", false}},
		{&service.SerTask{3, "333333", false}},
	}

	tests := []struct {
		name  string
		input *service.SerFilter
		want  []*service.SerTask
		//wantErr bool
	}{
		{"1 (1) - ", &service.SerFilter{0, []int{1}, " ", false}, []*service.SerTask{&service.SerTask{1, "111111", false}}},
		{"2 (2) - ", &service.SerFilter{0, []int{2}, " ", false}, []*service.SerTask{&service.SerTask{2, "222222", false}}},
		{"3 (3) - ", &service.SerFilter{0, []int{3}, " ", false}, []*service.SerTask{&service.SerTask{3, "333333", false}}},
		{"4 (1,2) - ", &service.SerFilter{0, []int{1, 2}, " ", false}, []*service.SerTask{&service.SerTask{1, "111111", false},
			&service.SerTask{2, "222222", false}}},
		{"5 (1,3) - ", &service.SerFilter{0, []int{1, 3}, " ", false}, []*service.SerTask{&service.SerTask{1, "111111", false},
			&service.SerTask{3, "333333", false}}},
		{"6 (2,3) - ", &service.SerFilter{0, []int{2, 3}, " ", false}, []*service.SerTask{&service.SerTask{2, "222222", false},
			&service.SerTask{3, "333333", false}}},
		{"7 (1,2,3) - ", &service.SerFilter{0, []int{1, 2, 3}, " ", false}, []*service.SerTask{&service.SerTask{1, "111111", false},
			&service.SerTask{2, "222222", false}, &service.SerTask{3, "333333", false}}},
		{"8 ( ) - ", &service.SerFilter{0, []int{}, " ", false}, []*service.SerTask{&service.SerTask{1, "111111", false},
			&service.SerTask{2, "222222", false}, &service.SerTask{3, "333333", false}}},
	}

	s := InitTest()

	//создание тасков
	//InitCreate(s)
	for _, task := range defCreate {
		err := s.Create(task.create)
		if err != nil {
			fmt.Printf("Create() error = %v", err)
		}
	}

	//проверка Read
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			readTask := s.Read(tc.input)

			/* 	if (readTask == tc.want)  { // если ошибка не нил , и не ждем ошибку , то ...
				t.Fatalf("что-то пошло не так!")
				return
			} */

			if !reflect.DeepEqual(tc.want, readTask) { //если нет ошибки , то сравниваем значения
				t.Fatalf("expected: %v, got: %v", tc.want, readTask)
			}
		})
	}

}
