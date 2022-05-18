package web

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	taskService "github.com/antoha2/todo/service/taskService"
)

func (wImpl *webImpl) Start() error {

	wImpl.server = &http.Server{Addr: ":8181"}

	mux := http.NewServeMux()
	//mux.HandleFunc("/api", wImpl.UserIdentify())

	mux.HandleFunc("/api/create", wImpl.UserIdentify(wImpl.handlerCreate))
	mux.HandleFunc("/api/read", wImpl.UserIdentify(wImpl.handlerRead))
	mux.HandleFunc("/api/delete", wImpl.UserIdentify(wImpl.handlerDelete))
	mux.HandleFunc("/api/update", wImpl.UserIdentify(wImpl.handlerUpdate))

	mux.HandleFunc("/auth/sign-up", wImpl.signUp)
	mux.HandleFunc("/auth/sign-in", wImpl.signIn)
	mux.HandleFunc("/auth/deleteUser", wImpl.deleteUser)

	log.Printf("Запуск веб-сервера на http://127.0.0.1:%s\n", wImpl.server.Addr) //:8181
	http.ListenAndServe(wImpl.server.Addr, mux)

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
		fmt.Println("can't unmarshal: ", err.Error())
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

	ctx := r.Context()
	userId, ok := FromContext(ctx)
	if !ok {
		newErr := "UserId не найден"
		log.Println(newErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(newErr))
		return
	}

	task := new(taskService.SerTask)
	err := wImpl.Decoder(r, task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	task.UserId = userId
	err = wImpl.taskService.Create(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	str := fmt.Sprintf("создание задачи выполнено. id-(%v) text-(%v) isDone-(%v)", task.Id, task.Text, task.IsDone)
	json, err := json.Marshal(str)

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

	ctx := r.Context()
	userId, ok := FromContext(ctx)
	if !ok {
		newErr := "UserId не найден"
		log.Println(newErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(newErr))
		return
	}

	readIds := new(taskService.SerFilter)

	err := wImpl.DecoderFilter(r, readIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	readIds.UserId = userId
	task := wImpl.taskService.Read(readIds)

	//str := fmt.Sprintf("чтение задачи выполнено - %v", task)
	for _, t := range task {
		str := fmt.Sprintf(" прочитана задача  id-(%v) text-(%v) isDone-(%v)   ", t.Id, t.Text, t.IsDone)
		json, err := json.Marshal(str)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(json)
	}
}

//обработчик Delete
func (wImpl *webImpl) handlerDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	ctx := r.Context()
	userId, ok := FromContext(ctx)
	if !ok {
		newErr := "UserId не найден"
		log.Println(newErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(newErr))
		return
	}

	delId := new(taskService.SerTask)
	delId.UserId = userId

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

	str := fmt.Sprintf("удаление задачи выполнено (id - %v)", delId.Id)
	json, err := json.Marshal(str)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(json)
}

//обработчик update
func (wImpl *webImpl) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	ctx := r.Context()
	userId, ok := FromContext(ctx)
	if !ok {
		newErr := "UserId не найден"
		log.Println(newErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(newErr))
		return
	}

	task := new(taskService.SerTask)

	err := wImpl.Decoder(r, task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	task.UserId = userId
	err = wImpl.taskService.Update(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	str := fmt.Sprintf("изменение задачи выполнено (id - %v)", task.Id)
	json, err := json.Marshal(str)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(json)
}
