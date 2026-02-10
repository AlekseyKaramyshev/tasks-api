package main

import (
	"log"
	"net/http"

	"github.com/AlekseyKaramyshev/tasks-api/internal/handlers"
	"github.com/AlekseyKaramyshev/tasks-api/internal/storage/memory"
)

func main() {
	storage := memory.New()

	h := handlers.New(storage)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", h.TasksCollection)
	mux.HandleFunc("/tasks/", h.TaskItem)

	log.Println("server listening on localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err)
	}
}
