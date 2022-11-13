package main

import (
	"log"
	"permata-aksesoris/apps/api/databases"
	"permata-aksesoris/apps/api/modules/categories"
	"permata-aksesoris/apps/api/modules/products"
	"permata-aksesoris/apps/api/modules/users"
)

func main() {
	con, err := databases.NewMainDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	if err := con.AutoMigrate(&categories.Category{}, &users.User{}, &products.Product{}, &products.ProductImage{}); err != nil {
		log.Fatal(err)
	}
}
