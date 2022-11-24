package modules

import (
	"net/http"
	"permata-aksesoris/apps/api/middlewares"
	"permata-aksesoris/apps/api/modules/categories"
	"permata-aksesoris/apps/api/modules/products"
	"permata-aksesoris/apps/api/modules/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to permata aksesoris API !"))
}

func NewApp(con *gorm.DB) http.HandlerFunc {
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

	routerWithMiddlewares := middlewares.NewLoggerMiddleware(middlewares.NewCorsMiddleware(router.ServeHTTP))

	return routerWithMiddlewares
}
