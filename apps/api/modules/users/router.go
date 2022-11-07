package users

import "github.com/gorilla/mux"

func NewRouter(handler Handler, router *mux.Router) {
	router.HandleFunc("/users/login", handler.login).Methods("POST")
}
