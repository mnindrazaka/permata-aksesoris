package main

import (
	"log"
	"permata-aksesoris/apps/api/databases"
)

func main() {
	con, err := databases.NewMainDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	if err := databases.Migrate(con); err != nil {
		log.Fatal(err)
	}
}
