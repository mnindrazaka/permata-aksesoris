package main

import (
	"log"
	"net/http"
	"os"
	"permata-aksesoris/apps/api/databases"
	"permata-aksesoris/apps/api/modules"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return
	}
	con, err := databases.NewMainDBConnection()
	if err != nil {
		log.Fatal(err)
		return
	}

	router := modules.NewApp(con)

	port := os.Getenv("PORT")
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
