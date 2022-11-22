package middlewares

import (
	"log"
	"net/http"
)

func NewLoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s, %s, %s\n", r.RemoteAddr, r.Method, r.URL)
		next(w, r)
	}
}
