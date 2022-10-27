package categories

import "github.com/gorilla/mux"

func NewRouter(handler Handler, router *mux.Router) {
	router.HandleFunc("/categories", handler.getCategories).Methods("GET")
	router.HandleFunc("/categories", handler.createCategory).Methods("POST")
}
