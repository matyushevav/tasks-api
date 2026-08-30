package main

import (
	"log"
	"net/http"

	"tasks-api/internal/handlers"
	"tasks-api/internal/middleware"
	"tasks-api/internal/storage"
)

func main() {
	store := storage.NewMemoryStorage()
	h := handlers.New(store)

	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", middleware.LoggingMiddleware(h.TasksCollection))
	mux.HandleFunc("/tasks/", middleware.LoggingMiddleware(h.TaskItem))

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
