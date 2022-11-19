package main

import (
	"log"
	"net/http"
	"permata-aksesoris/apps/api/databases"
	"permata-aksesoris/apps/api/modules"
)

func main() {
	con, err := databases.NewMainDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := modules.NewApp(con)

	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
