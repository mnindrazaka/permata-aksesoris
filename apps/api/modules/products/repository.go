package products

import "gorm.io/gorm"

type repository struct {
	con *gorm.DB
}

type Repository interface {
	getProducts() ([]Product, error)
}

func NewRepository(con *gorm.DB) Repository {
	return repository{con}
}

func (repository repository) getProducts() ([]Product, error) {
	var products []Product
	result := repository.con.Model(Product{}).Preload("Images").Joins("Category").Find(&products)
	return products, result.Error
}
