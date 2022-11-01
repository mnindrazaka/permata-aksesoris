package products

import (
	"permata-aksesoris/apps/api/utils"

	"gorm.io/gorm"
)

type repository struct {
	con *gorm.DB
}

type Repository interface {
	getProducts() ([]Product, error)
	getProductDetail(string) (Product, error)
	createProduct(Product) (Product, error)
	updateProduct(string, Product) (Product, error)
	deleteProduct(string) error
}

func NewRepository(con *gorm.DB) Repository {
	return repository{con}
}

func (repository repository) getProducts() ([]Product, error) {
	var products []Product
	result := repository.con.Model(Product{}).Preload("Category").Preload("Images").Find(&products)
	return products, result.Error
}

func (repository repository) getProductDetail(serial string) (Product, error) {
	var product Product
	result := repository.con.Model(Product{}).Preload("Category").Preload("Images").First(&product, Product{Serial: serial})
	return product, result.Error
}

func (repository repository) createProduct(product Product) (Product, error) {
	for index := range product.Images {
		product.Images[index].Serial = utils.CreateSerial("IMG")
	}

	product.Serial = utils.CreateSerial("PRD")
	result := repository.con.Model(Product{}).Create(product)
	return product, result.Error
}

func (repository repository) updateProduct(serial string, product Product) (Product, error) {
	product.Serial = serial
	result := repository.con.Session(&gorm.Session{FullSaveAssociations: true}).Model(&product).Updates(product)
	return product, result.Error
}

func (repository repository) deleteProduct(serial string) error {
	result := repository.con.Model(Product{}).Delete(Product{Serial: serial})
	return result.Error
}
