package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware - логирует метод и путь запроса
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}
