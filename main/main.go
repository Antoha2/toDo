package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/antoha2/todo/config"
	authService "github.com/antoha2/todo/service/authService"
	authRepository "github.com/antoha2/todo/service/authService/authRepository"
	taskService "github.com/antoha2/todo/service/taskService"
	taskRepository "github.com/antoha2/todo/service/taskService/taskRepository"
	web "github.com/antoha2/todo/transport/web"
	//authWeb "github.com/antoha2/todo/transport/web/authWeb"
)

func main() {

	Run()

}

func initDb(cfg *config.Config) (*sqlx.DB, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Dbname,
		cfg.DB.Sslmode,
	)

	// Prep config
	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf(" failed to parse config: %v", err)
	}

	// Make connections
	dbx, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf(" failed to create connection db: %v", err)
	}

	err = dbx.Ping()
	if err != nil {
		return nil, fmt.Errorf(" error to ping connection pool: %v", err)
	}
	log.Printf("Запуск базы данных на http://127.0.0.1:%d\n", cfg.DB.Port)
	return dbx, nil
}

func Run() {

	cfg := config.GetConfig()
	dbx, err := initDb(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	taskRep := taskRepository.NewTaskRepository(dbx)
	authRep := authRepository.NewAuthPostgres(dbx)

	taskSer := taskService.NewTaskService(taskRep)
	authSer := authService.NewAuthService(authRep)

	Tran := web.NewWeb(taskSer, *authSer)

	go Tran.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	Tran.Stop()

}
