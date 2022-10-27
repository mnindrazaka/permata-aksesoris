package categories

import (
	"permata-aksesoris/apps/api/utils"

	"gorm.io/gorm"
)

type repository struct {
	con *gorm.DB
}

type Repository interface {
	getCategories() ([]Category, error)
	createCategory(Category) (Category, error)
}

func NewRepository(con *gorm.DB) Repository {
	return repository{con}
}

func (repo repository) getCategories() ([]Category, error) {
	var categories []Category
	result := repo.con.Model(Category{}).Find(&categories)
	return categories, result.Error
}

func (repo repository) createCategory(category Category) (Category, error) {
	category.Serial = utils.CreateSerial("CAT")
	result := repo.con.Model(Category{}).Create(category)
	return category, result.Error
}
