package databases

import (
	"permata-aksesoris/apps/api/modules/categories"
	"permata-aksesoris/apps/api/modules/products"
	"permata-aksesoris/apps/api/modules/users"

	"gorm.io/gorm"
)

func Migrate(con *gorm.DB) error {
	return con.AutoMigrate(&categories.Category{}, &users.User{}, &products.Product{}, &products.ProductImage{})
}
