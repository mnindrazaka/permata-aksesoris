package categories

import (
	"gorm.io/gorm"
)

type repository struct {
	con *gorm.DB
}

type Repository interface {
	getCategories() ([]Category, error)
}

func NewRepository(con *gorm.DB) Repository {
	return repository{con}
}

func (repo repository) getCategories() ([]Category, error) {
	var categories []Category
	result := repo.con.Table("categories").Find(&categories)
	return categories, result.Error
}
