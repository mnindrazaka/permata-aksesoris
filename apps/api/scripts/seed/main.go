package main

import (
	"log"
	"permata-aksesoris/apps/api/databases"
)

func main() {
	con, err := databases.NewMainDBConnection()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := databases.Seed(con); err != nil {
		log.Fatal(err)
	}
}
