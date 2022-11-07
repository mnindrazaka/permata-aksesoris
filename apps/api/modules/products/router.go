package products

import (
	"permata-aksesoris/apps/api/middlewares"

	"github.com/gorilla/mux"
)

func NewRouter(handler Handler, router *mux.Router) {
	router.HandleFunc("/products", handler.getProducts).Methods("GET")
	router.HandleFunc("/products/{serial}", handler.getProductDetail).Methods("GET")
	router.HandleFunc("/products", middlewares.NewAuthenticateMiddleware(handler.createProduct)).Methods("POST")
	router.HandleFunc("/products/{serial}", middlewares.NewAuthenticateMiddleware(handler.updateProduct)).Methods("PUT")
	router.HandleFunc("/products/{serial}", middlewares.NewAuthenticateMiddleware(handler.deleteProduct)).Methods("DELETE")

	router.HandleFunc("/products/{productSerial}/images", middlewares.NewAuthenticateMiddleware(handler.createProductImage)).Methods("POST")
	router.HandleFunc("/products/{productSerial}/images/{imageSerial}", middlewares.NewAuthenticateMiddleware(handler.deleteProductImage)).Methods("DELETE")
}
