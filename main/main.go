package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/antoha2/todo/repository"
	"github.com/antoha2/todo/service"
	"github.com/antoha2/todo/transport/web"
)

func main() {

	Run()

}

func Run() {

	rep := repository.NewMap()
	ser := service.New(rep)
	tran := web.New(ser)

	go tran.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	tran.Stop()

}
