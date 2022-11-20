package main

import (
	"log"
	"net/http"
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

	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
