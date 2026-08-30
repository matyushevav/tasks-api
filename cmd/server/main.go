package main

import (
    "encoding/json"
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
    
    // Обработчик для всех остальных маршрутов
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Проверяем, что это не /tasks и не /tasks/ (это уже обработано выше)
        if r.URL.Path != "/" {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusNotFound)
            json.NewEncoder(w).Encode(map[string]string{"error": "Endpoint not found"})
            return
        }
        // Для корневого пути можно вернуть список эндпоинтов
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Tasks API. Available endpoints: GET/POST /tasks, GET/PUT/DELETE /tasks/{id}",
        })
    })

    log.Println("Server listening on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}