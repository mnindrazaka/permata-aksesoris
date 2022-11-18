package tests

import (
	"net/http"
	"permata-aksesoris/apps/api/databases"
	"permata-aksesoris/apps/api/middlewares"
	"permata-aksesoris/apps/api/modules/categories"
	"permata-aksesoris/apps/api/modules/products"
	"permata-aksesoris/apps/api/modules/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func BeforeEach() (*gorm.DB, http.HandlerFunc, error) {
	con, err := databases.NewTestDBConnection()
	if err != nil {
		return nil, nil, err
	}

	if err := databases.Migrate(con); err != nil {
		return nil, nil, err
	}

	if err := databases.Unseed(con); err != nil {
		return nil, nil, err
	}

	if err := databases.Seed(con); err != nil {
		return nil, nil, err
	}

	router := mux.NewRouter().StrictSlash(true)

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

	return con, routerWithCors, nil
}

func AfterEach(con *gorm.DB) error {
	return databases.Unseed(con)
}
