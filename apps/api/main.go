package main

import (
	"log"
	"net/http"
	"permata-aksesoris/apps/api/middlewares"
	"permata-aksesoris/apps/api/modules/categories"
	"permata-aksesoris/apps/api/modules/products"
	"permata-aksesoris/apps/api/modules/users"
	"permata-aksesoris/apps/api/utils"

	"github.com/gorilla/mux"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to permata aksesoris API"))
}

func main() {
	con, err := utils.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleHome)

	categoryRepository := categories.NewRepository(con)
	categoryUsecase := categories.NewUsecase(categoryRepository)
	categoryHandler := categories.NewHandler(categoryUsecase)
	categories.NewRouter(categoryHandler, router)

	productRepository := products.NewRepository(con)
	productUsecase := products.NewUsecase(productRepository)
	productHandler := products.NewHandler(productUsecase)
	products.NewRouter(productHandler, router)

	userRepository := users.NewRepository(con)
	userUsecase := users.NewUsecase(userRepository)
	userHandler := users.NewHandler(userUsecase)
	users.NewRouter(userHandler, router)

	routerWithCors := middlewares.NewCorsMiddleware(router.ServeHTTP)

	err = http.ListenAndServe(":3000", routerWithCors)
	if err != nil {
		log.Fatal(err)
	}
}
