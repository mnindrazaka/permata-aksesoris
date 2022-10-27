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
	updateCategory(string, Category) (Category, error)
	deleteCategory(string) error
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

func (repo repository) updateCategory(serial string, category Category) (Category, error) {
	category.Serial = serial
	result := repo.con.Model(&category).Updates(category)
	return category, result.Error
}

func (repo repository) deleteCategory(serial string) error {
	result := repo.con.Model(Category{}).Delete(Category{Serial: serial})
	return result.Error
}
