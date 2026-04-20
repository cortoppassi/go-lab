package main

import (
	"log"
	"net/http"

	"go-lab/internal/httpapi"
	"go-lab/internal/task"
)

func main() {
	repository := task.NewMemoryRepository()
	service := task.NewService(repository)
	handler := task.NewHandler(service)

	_, err := service.CreateTask(task.CreateTaskInput{
		Title:       "Estudar Go",
		Description: "Organizar um CRUD em camadas",
		Done:        false,
	})
	if err != nil {
		log.Fatal(err)
	}

	router := httpapi.NewRouter(handler.RegisterRoutes)

	log.Println("Servidor iniciado em http://localhost:8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
