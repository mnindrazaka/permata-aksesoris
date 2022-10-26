package products

import "permata-aksesoris/apps/api/modules/categories"

type Product struct {
	Serial         string              `json:"serial" gorm:"primaryKey"`
	Title          string              `json:"title"`
	Slug           string              `json:"slug"`
	Thumbnail      string              `json:"thumbnail"`
	Description    string              `json:"description"`
	CategorySerial string              `json:"-"`
	Category       categories.Category `json:"category" gorm:"foreignKey:CategorySerial"`
	Images         []ProductImages     `json:"images" gorm:"foreignKey:ProductSerial"`
}

type ProductImages struct {
	Serial        string `json:"serial" gorm:"primaryKey"`
	Url           string `json:"url"`
	ProductSerial string `json:"-"`
}
