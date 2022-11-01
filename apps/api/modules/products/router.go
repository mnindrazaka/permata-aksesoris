package products

import "github.com/gorilla/mux"

func NewRouter(handler Handler, router *mux.Router) {
	router.HandleFunc("/products", handler.getProducts).Methods("GET")
	router.HandleFunc("/products/{serial}", handler.getProductDetail).Methods("GET")
	router.HandleFunc("/products", handler.createProduct).Methods("POST")
	router.HandleFunc("/products/{serial}", handler.updateProduct).Methods("PUT")
	router.HandleFunc("/products/{serial}", handler.deleteProduct).Methods("DELETE")
}
