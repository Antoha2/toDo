package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/antoha2/todo/repository"
	"github.com/antoha2/todo/service"
)

type defCreate []struct {
	create *service.SerTask
}

func InitTest() service.Service {
	rep := repository.NewMap()
	ser := service.New(rep)
	return ser
}

func InitCreate(s service.Service) {
	defCreate := defCreate{
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
func TestCreateTask(t *testing.T) {
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
}

//test Read
func TestReadTask(t *testing.T) {

	tests := []struct {
		name    string
		input   *service.SerFilter
		want    []*service.SerTask
		wantErr bool
	}{
		{"1 (1) - ", &service.SerFilter{0, []int{1}, " ", false}, []*service.SerTask{&service.SerTask{1, "111111", false}}, false},
		{"2 (2) - ", &service.SerFilter{0, []int{2}, " ", false}, []*service.SerTask{&service.SerTask{2, "222222", false}}, false},
		{"3 (3) - ", &service.SerFilter{0, []int{3}, " ", false}, []*service.SerTask{&service.SerTask{3, "333333", false}}, false},
		{"4 (1,2) - ", &service.SerFilter{0, []int{1, 2}, " ", false}, []*service.SerTask{&service.SerTask{1, "111111", false},
			&service.SerTask{2, "222222", false}}, false},
		{"5 (1,3) - ", &service.SerFilter{0, []int{1, 3}, " ", false}, []*service.SerTask{&service.SerTask{1, "111111", false},
			&service.SerTask{3, "333333", false}}, false},
		{"6 (2,3) - ", &service.SerFilter{0, []int{2, 3}, " ", false}, []*service.SerTask{&service.SerTask{2, "222222", false},
			&service.SerTask{3, "333333", false}}, false},
		{"7 (1,2,3) - ", &service.SerFilter{0, []int{1, 2, 3}, " ", false}, []*service.SerTask{&service.SerTask{1, "111111", false},
			&service.SerTask{2, "222222", false}, &service.SerTask{3, "333333", false}}, false},
		{"8 ( ) - ", &service.SerFilter{0, []int{}, " ", false}, []*service.SerTask{&service.SerTask{1, "111111", false},
			&service.SerTask{2, "222222", false}, &service.SerTask{3, "333333", false}}, false},
		{"9 (42) - ", &service.SerFilter{0, []int{42}, " ", false}, []*service.SerTask{}, true},
	}

	s := InitTest()

	//создание тасков
	InitCreate(s)

	//проверка Read
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			readTask := s.Read(tc.input)
			if !reflect.DeepEqual(tc.want, readTask) { //если нет ошибки , то сравниваем значения

				t.Fatalf("expected: %v, got: %v", tc.want, readTask)
			}
		})
	}
}

//test Delete
func TestDeleteTask(t *testing.T) {

	tests := []struct {
		name    string
		input   *service.SerTask
		want    []*service.SerTask
		wantErr bool
	}{
		{"удаляю первую задачу", &service.SerTask{3, "333333", false}, []*service.SerTask{&service.SerTask{1, "111111", false}, &service.SerTask{2, "222222", false}}, false},
		{"удаляю вторую задачу", &service.SerTask{2, "222222", false}, []*service.SerTask{&service.SerTask{1, "111111", false}}, false},
		{"удаляю третью задачу", &service.SerTask{1, "111111", false}, []*service.SerTask{}, false},
		{"удаляю четвертую задачу", &service.SerTask{4, "444444", false}, []*service.SerTask{}, true},
	}

	s := InitTest()

	//создание тасков
	InitCreate(s)

	//проверка Delete
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := s.Delete(tc.input)
			if err != nil {
				t.Fatalf("Update() error = %v", err)
				return
			}

			wantIds := make([]int, len(tc.want))
			for index := 0; index < len(tc.want); index++ {

				wantIds[index] = tc.want[index].Id
			}

			readTask := s.Read(&service.SerFilter{Ids: wantIds})

			if (err != nil) != tc.wantErr { // если ошибка не нил , и не ждем ошибку , то ...
				t.Fatalf("Update() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if len(readTask) == 0 {
				return
			}
			if !reflect.DeepEqual(tc.want, readTask) { //если нет ошибки , то сравниваем значения
				t.Fatalf("expected: %v, got: %v", tc.want, readTask)
			}
		})
	}

}

//test Update
func TestUpdateTask(t *testing.T) {

	tests := []struct {
		name    string
		input   *service.SerTask
		want    *service.SerTask
		wantErr bool
	}{
		{"изменяю первую задачу", &service.SerTask{1, "1-1-1-1-1-1", false}, &service.SerTask{1, "1-1-1-1-1-1", false}, false},
		{"изменяю вторую задачу", &service.SerTask{2, "2-2-2-2-2-2", true}, &service.SerTask{2, "2-2-2-2-2-2", true}, false},
		{"изменяю третью задачу", &service.SerTask{3, "3-3-3-3-3-3", true}, &service.SerTask{3, "3-3-3-3-3-3", true}, false},
		{"изменяю четвертую задачу", &service.SerTask{4, "4-4-4-4-4-4", true}, &service.SerTask{4, "4-4-4-4-4-4", true}, true},
	}

	s := InitTest()

	//создание тасков
	InitCreate(s)

	//проверка Update
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := s.Update(tc.input)
			readTask := s.Read(&service.SerFilter{Ids: []int{tc.want.Id}})

			if (err != nil) != tc.wantErr { // если ошибка не нил , и не ждем ошибку , то ...
				t.Fatalf("Update() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if len(readTask) == 0 {
				return
			}
			if !reflect.DeepEqual(*tc.want, *readTask[0]) { //если нет ошибки , то сравниваем значения
				t.Fatalf("expected: %v, got: %v", *tc.want, *tc.input)
			}
		})
	}
}
