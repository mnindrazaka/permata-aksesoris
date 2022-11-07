package categories

import (
	"permata-aksesoris/apps/api/middlewares"

	"github.com/gorilla/mux"
)

func NewRouter(handler Handler, router *mux.Router) {
	router.HandleFunc("/categories", handler.getCategories).Methods("GET")
	router.HandleFunc("/categories", middlewares.NewAuthenticateMiddleware(handler.createCategory)).Methods("POST")
	router.HandleFunc("/categories/{serial}", middlewares.NewAuthenticateMiddleware(handler.updateCategory)).Methods("PUT")
	router.HandleFunc("/categories/{serial}", middlewares.NewAuthenticateMiddleware(handler.deleteCategory)).Methods("DELETE")
}
