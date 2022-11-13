package main

import (
	"log"
	"permata-aksesoris/apps/api/databases"
	"permata-aksesoris/apps/api/modules/categories"
	"permata-aksesoris/apps/api/modules/products"
	"permata-aksesoris/apps/api/modules/users"
	"permata-aksesoris/apps/api/utils"
)

var usersData = []users.User{
	{Serial: utils.CreateSerial("USR"), Email: "admin@gmail.com", Password: "admin"},
	{Serial: utils.CreateSerial("USR"), Email: "user@gmail.com", Password: "user"},
}

var categoriesData = []categories.Category{
	{Serial: utils.CreateSerial("CAT"), Title: "Kacamata", Slug: "kacamata"},
	{Serial: utils.CreateSerial("CAT"), Title: "Jam Tangan", Slug: "jam-tangan"},
}

var productsData = []products.Product{
	{Serial: utils.CreateSerial("PRD"), Title: "Kacamata Gaul", Slug: "kacamata-gaul", Description: "kacamata gaul sekali", Thumbnail: "https://img.freepik.com/free-photo/eyeglasses-wear_1203-2604.jpg", CategorySerial: categoriesData[0].Serial},
	{Serial: utils.CreateSerial("PRD"), Title: "Jam Tangan Gaul", Slug: "jam-tangan-gaul", Description: "jam tangan gaul sekali", Thumbnail: "https://img.freepik.com/free-photo/businessman-checking-time_1357-97.jpg", CategorySerial: categoriesData[1].Serial},
}

var productImagesData = []products.ProductImage{
	{Serial: utils.CreateSerial("IMG"), Url: "https://img.freepik.com/free-photo/eyeglasses-wear_1203-2604.jpg", ProductSerial: productsData[0].Serial},
	{Serial: utils.CreateSerial("IMG"), Url: "https://img.freepik.com/free-photo/businessman-checking-time_1357-97.jpg", ProductSerial: productsData[1].Serial},
}

func main() {
	con, err := databases.NewTestDBConnection()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, user := range usersData {
		if err := con.Model(users.User{}).Create(user).Error; err != nil {
			log.Fatal(err)
		}
	}

	for _, category := range categoriesData {
		if err := con.Model(categories.Category{}).Create(category).Error; err != nil {
			log.Fatal(err)
		}
	}

	for _, product := range productsData {
		if err := con.Model(products.Product{}).Create(product).Error; err != nil {
			log.Fatal(err)
		}
	}

	for _, productImage := range productImagesData {
		if err := con.Model(products.ProductImage{}).Create(productImage).Error; err != nil {
			log.Fatal(err)
		}
	}
}
