package databases

import (
	"permata-aksesoris/apps/api/modules/categories"
	"permata-aksesoris/apps/api/modules/products"
	"permata-aksesoris/apps/api/modules/users"

	"gorm.io/gorm"
)

var UsersData = []users.User{
	{Serial: "USR-1", Email: "admin@gmail.com", Password: "admin"},
	{Serial: "USR-2", Email: "user@gmail.com", Password: "user"},
}

var CategoriesData = []categories.Category{
	{Serial: "CAT-1", Title: "Kacamata", Slug: "kacamata"},
	{Serial: "CAT-2", Title: "Jam Tangan", Slug: "jam-tangan"},
}

var ProductsData = []products.Product{
	{Serial: "PRD-1", Title: "Kacamata Gaul", Slug: "kacamata-gaul", Description: "kacamata gaul sekali", Thumbnail: "https://img.freepik.com/free-photo/eyeglasses-wear_1203-2604.jpg", CategorySerial: CategoriesData[0].Serial},
	{Serial: "PRD-2", Title: "Jam Tangan Gaul", Slug: "jam-tangan-gaul", Description: "jam tangan gaul sekali", Thumbnail: "https://img.freepik.com/free-photo/businessman-checking-time_1357-97.jpg", CategorySerial: CategoriesData[1].Serial},
}

var ProductImagesData = []products.ProductImage{
	{Serial: "IMG-1", Url: "https://img.freepik.com/free-photo/eyeglasses-wear_1203-2604.jpg", ProductSerial: ProductsData[0].Serial},
	{Serial: "IMG-2", Url: "https://img.freepik.com/free-photo/businessman-checking-time_1357-97.jpg", ProductSerial: ProductsData[1].Serial},
}

func Seed(con *gorm.DB) error {
	for _, user := range UsersData {
		if err := con.Model(users.User{}).Create(user).Error; err != nil {
			return err
		}
	}

	for _, category := range CategoriesData {
		if err := con.Model(categories.Category{}).Create(category).Error; err != nil {
			return err
		}
	}

	for _, product := range ProductsData {
		if err := con.Model(products.Product{}).Create(product).Error; err != nil {
			return err
		}
	}

	for _, productImage := range ProductImagesData {
		if err := con.Model(products.ProductImage{}).Create(productImage).Error; err != nil {
			return err
		}
	}

	return nil
}

func Unseed(con *gorm.DB) error {
	if err := con.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(products.ProductImage{}).Delete(products.ProductImage{}).Error; err != nil {
		return err
	}
	if err := con.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(products.Product{}).Delete(products.Product{}).Error; err != nil {
		return err
	}
	if err := con.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(categories.Category{}).Delete(categories.Category{}).Error; err != nil {
		return err
	}
	if err := con.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(users.User{}).Delete(users.User{}).Error; err != nil {
		return err
	}
	return nil
}
