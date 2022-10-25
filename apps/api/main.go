package main

import (
	"log"
	"net/http"
	"permata-aksesoris/apps/api/modules/categories"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to permata aksesoris API"))
}

func main() {
	dsn := "root:roottoor@tcp(127.0.0.1:3306)/permata_aksesoris?charset=utf8mb4&parseTime=True&loc=Local"

	con, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleHome)

	categoryRepository := categories.NewRepository(con)
	categoryUsecase := categories.NewUsecase(categoryRepository)
	categoryHandler := categories.NewHandler(categoryUsecase)
	categories.NewRouter(categoryHandler, router)

	http.ListenAndServe(":3000", router)
}
