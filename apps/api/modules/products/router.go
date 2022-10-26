package products

import "github.com/gorilla/mux"

func NewRouter(handler Handler, router *mux.Router) {
	router.HandleFunc("/products", handler.getProducts).Methods("GET")
}
