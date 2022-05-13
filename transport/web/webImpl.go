package web

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	taskService "github.com/antoha2/todo/service"
	authService "github.com/antoha2/todo/service/authService"
)

func NewTaskWeb(service taskService.TodolistServ) *webImpl {
	return &webImpl{
		taskService: service, //??????
		//authService: authService,
	}
}

func NewAuthWeb(authService authService.Authorization) *authWebImpl {
	return &authWebImpl{
		authService: authService,
	}
}

func (wImpl *webImpl) StartTask() error {

	wImpl.server = &http.Server{Addr: ":8181"}

	mux := http.NewServeMux()
	mux.HandleFunc("/create", wImpl.handlerCreate)
	mux.HandleFunc("/read", wImpl.handlerRead)
	mux.HandleFunc("/delete", wImpl.handlerDelete)
	mux.HandleFunc("/update", wImpl.handlerUpdate)

	log.Println("Запуск веб-сервера на http://127.0.0.1:8181")
	http.ListenAndServe(":8181", mux)

	return nil
}

func (wImpl *webImpl) Stop() {

	if err := wImpl.server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}

//декодеры JSON
func (wImpl *webImpl) Decoder(r *http.Request, task *taskService.SerTask) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, task)
	if err != nil {
		fmt.Println("can't unmarshal !!!!!: ", err.Error())
		return err
	}
	return nil
}

func (wImpl *webImpl) DecoderFilter(r *http.Request, task *taskService.SerFilter) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, task)
	if err != nil {
		fmt.Println("can't unmarshal: ", err.Error())
		return err
	}
	return nil
}

//обработчик Сreate
func (wImpl *webImpl) handlerCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	task := new(taskService.SerTask)

	err := wImpl.Decoder(r, task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = wImpl.taskService.Create(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(task)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(json)

}

//обработчик Read
func (wImpl *webImpl) handlerRead(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	readIds := new(taskService.SerFilter)

	err := wImpl.DecoderFilter(r, readIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	task := wImpl.taskService.Read(readIds)
	json, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(json)

}

//обработчик Delete
func (wImpl *webImpl) handlerDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	delId := new(taskService.SerTask)

	err := wImpl.Decoder(r, delId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = wImpl.taskService.Delete(delId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

//обработчик update
func (wImpl *webImpl) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	task := new(taskService.SerTask)

	err := wImpl.Decoder(r, task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = wImpl.taskService.Update(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(task)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(json)

}
