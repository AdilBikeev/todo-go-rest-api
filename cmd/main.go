package main

import (
	todo_app "github.com/adcoder/todo-app"
	"github.com/adcoder/todo-app/pkg/handler"
	"github.com/adcoder/todo-app/pkg/repository"
	"github.com/adcoder/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run("8008", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}