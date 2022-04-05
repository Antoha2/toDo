package main

import (
	"os"
	"os/signal"
	"syscall"

	//"./etodo/repository"
	//"github.com/staszigzag/sandbox/config"
	"github.com/antoha2/todo/repository"
	"github.com/antoha2/todo/service"
	"github.com/antoha2/todo/transport/web"
	/*"github.com/staszigzag/sandbox/repository2"
	"github.com/staszigzag/sandbox/service"
	"github.com/staszigzag/sandbox/transport/web" */)

func main() {

	Run()

}

func Run() {

	rep := repository.New()
	ser := service.New(rep)
	tran := web.New(ser)

	go tran.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	tran.Stop()

}
