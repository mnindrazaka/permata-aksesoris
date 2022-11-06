package middlewares

import "net/http"

type cors struct {
	nextHandler http.Handler
}

type Cors interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func NewCors(nextHandler http.Handler) Cors {
	return cors{nextHandler}
}

func (cors cors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
	cors.nextHandler.ServeHTTP(w, r)
}
