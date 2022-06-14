package web

import (
	"context"
	"encoding/json"

	"log"
	"net/http"

	authEndpoints "github.com/antoha2/todo/transport/web/endpoints/authEndpoints"
	taskEndpoints "github.com/antoha2/todo/transport/web/endpoints/taskEndpoints"

	//	authEndpoints "github.com/antoha2/todo/transport/web/endpoints
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func (wImpl *webImpl) Start() error {

	//options provided by the Go kit to facilitate error control
	TaskOptions := []httptransport.ServerOption{
		httptransport.ServerBefore(wImpl.UserIdentify),
	}

	AuthOptions := []httptransport.ServerOption{
		//httptransport.ServerBefore(wImpl.UserIdentify),
	}

	CreateHandler := httptransport.NewServer(
		taskEndpoints.MakeCreateEndpoint(wImpl.taskService), //use the endpoint
		decodeMakeCreateRequest,                             //converts the parameters received via the request body into the struct expected by the endpoint
		encodeResponse,
		TaskOptions...,
	)

	ReadHandler := httptransport.NewServer(
		taskEndpoints.MakeReadEndpoint(wImpl.taskService),
		decodeMakeReadRequest,
		encodeResponse,
		TaskOptions...,
	)

	UpdateHandler := httptransport.NewServer(
		taskEndpoints.MakeUpdateEndpoint(wImpl.taskService),
		decodeMakeUpdateRequest,
		encodeResponse,
		TaskOptions...,
	)

	DeleteHandler := httptransport.NewServer(
		taskEndpoints.MakeDeleteEndpoint(wImpl.taskService),
		decodeMakeDeleteRequest,
		encodeResponse,
		TaskOptions...,
	)
	signInHandler := httptransport.NewServer(
		authEndpoints.MakeSignInEndpoint(&wImpl.authService),
		decodeMakeSignInRequest,
		encodeResponse,
		AuthOptions...,
	)

	signUpAdminHandler := httptransport.NewServer(
		authEndpoints.MakeSignUpAdminEndpoint(&wImpl.authService),
		decodeMakeSignUpAdminRequest,
		encodeResponse,
		AuthOptions...,
	)

	signUpUserHandler := httptransport.NewServer(
		authEndpoints.MakeSignUpUserEndpoint(&wImpl.authService),
		decodeMakeSignUpUserRequest,
		encodeResponse,
		AuthOptions...,
	)

	deleteUserHandler := httptransport.NewServer(
		authEndpoints.MakeDeleteUserEndpoint(&wImpl.authService),
		decodeMakeDeleteUserRequest,
		encodeResponse,
		httptransport.ServerBefore(wImpl.UserIdentify),
	)

	updateUserHandler := httptransport.NewServer(
		authEndpoints.MakeUpdateUserEndpoint(&wImpl.authService),
		decodeMakeUpdateUserRequest,
		encodeResponse,
		httptransport.ServerBefore(wImpl.UserIdentify),
	)

	r := mux.NewRouter() //I'm using Gorilla Mux, but it could be any other library, or even the stdlib
	r.Methods("POST").Path("/api/create").Handler(CreateHandler)
	r.Methods("POST").Path("/api/read").Handler(ReadHandler)
	r.Methods("POST").Path("/api/update").Handler(UpdateHandler)
	r.Methods("POST").Path("/api/delete").Handler(DeleteHandler)

	r.Methods("POST").Path("/auth/sign-in").Handler(signInHandler)
	r.Methods("POST").Path("/auth/sign-up/admin").Handler(signUpAdminHandler)
	r.Methods("POST").Path("/auth/sign-up/user").Handler(signUpUserHandler)
	r.Methods("POST").Path("/auth/deleteUser").Handler(deleteUserHandler)
	r.Methods("POST").Path("/auth/updateUser").Handler(updateUserHandler)

	wImpl.server = &http.Server{Addr: ":8181"}
	log.Printf("Запуск веб-сервера на http://127.0.0.1%s\n", wImpl.server.Addr) //:8181

	if err := http.ListenAndServe(wImpl.server.Addr, r); err != nil {
		log.Println(err)
	}

	return nil

}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeMakeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request taskEndpoints.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	log.Println(r.Body)
	return request, nil
}

func decodeMakeReadRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request taskEndpoints.ReadRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request taskEndpoints.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request taskEndpoints.DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeSignInRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeSignUpAdminRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.SignUpAdminRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeSignUpUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.SignUpUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeDeleteUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.DeleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeMakeUpdateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request authEndpoints.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func (wImpl *webImpl) Stop() {

	if err := wImpl.server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}

//*/
/*
func decodeMakeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	log.Println("1")
	var request *endpoints.CreateRequest
	//request := new(endpoints.CreateRequest)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	//log.Println(body)
	//log.Println(request)

	err = json.Unmarshal(body, &request)
	if err != nil {
		fmt.Println("can't unmarshal: ", err.Error())
		return nil, err
	}

	log.Println(request)
	return request, nil
} //*/

/*
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


/*func (wImpl *webImpl) Start() error {

	wImpl.server = &http.Server{Addr: ":8181"}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/create", wImpl.UserIdentify(wImpl.handlerCreate))
	mux.HandleFunc("/api/read", wImpl.UserIdentify(wImpl.handlerRead))
	mux.HandleFunc("/api/update", wImpl.UserIdentify(wImpl.handlerUpdate))
	mux.HandleFunc("/api/delete", wImpl.UserIdentify(wImpl.handlerDelete))

	mux.HandleFunc("/auth/sign-up/admin", wImpl.signUpAdmin)
	mux.HandleFunc("/auth/sign-up/user", wImpl.signUpUser)
	mux.HandleFunc("/auth/sign-in", wImpl.signIn)
	mux.HandleFunc("/auth/deleteUser", wImpl.UserIdentify(wImpl.deleteUser))
	mux.HandleFunc("/auth/updateUser", wImpl.UserIdentify(wImpl.updateUser))

	log.Printf("Запуск веб-сервера на http://127.0.0.1%s\n", wImpl.server.Addr) //:8181

	if err := http.ListenAndServe(wImpl.server.Addr, mux); err != nil {
		log.Println(err)
	}

	return nil
} */
/*

decodeMakeCreateRequest


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

	userId, ok := r.Context().Value(etodo.USER_ID).(int)
	if !ok {
		newErr := "UserId не найден"
		//log.Println(newErr)
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
	err = wImpl.taskService.Create(r.Context(), task)
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

	userId, ok := r.Context().Value(etodo.USER_ID).(int)
	if !ok {
		newErr := "UserId не найден"
		//log.Println(newErr)
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
	task := wImpl.taskService.Read(r.Context(), readIds)
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

	userId, ok := r.Context().Value(etodo.USER_ID).(int)
	if !ok {
		newErr := "UserId не найден"
		//log.Println(newErr)
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

	err = wImpl.taskService.Delete(r.Context(), delId)
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

	userId, ok := r.Context().Value(etodo.USER_ID).(int)
	if !ok {
		newErr := "UserId не найден"
		//log.Println(newErr)
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

	err = wImpl.taskService.Update(r.Context(), task)
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
*/
