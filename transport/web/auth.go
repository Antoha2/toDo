package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	etodo "github.com/antoha2/todo"
	"github.com/gin-gonic/gin"
)

func (authWebImpl *authWebImpl) signUp(w http.ResponseWriter, r *http.Request) {

	user := new(etodo.User)

	if r.Method != http.MethodPost {
		return
	}

	err := authWebImpl.DecoderAuth(r, user)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = authWebImpl.authService.CreateUser(user)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(user)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(json)

}

func (h *webImpl) signIn(c *gin.Context) {

}

//Запуск веб-сервера Auth
func (authWebImpl *authWebImpl) StartAuth() error {

	authWebImpl.server = &http.Server{Addr: ":8180"}

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/sign-up", authWebImpl.signUp)

	log.Printf("Запуск веб-сервера на http://127.0.0.1:%s\n", authWebImpl.server.Addr)
	http.ListenAndServe(":8180", mux)

	return nil
}

//декодеры JSON
func (authWebImpl *authWebImpl) DecoderAuth(r *http.Request, user *etodo.User) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, user)
	if err != nil {
		fmt.Println("can't unmarshal: ", err.Error())
		return err
	}
	return nil
}
