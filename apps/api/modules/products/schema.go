package products

import "permata-aksesoris/apps/api/modules/categories"

type Product struct {
	Serial         string              `json:"serial" gorm:"primaryKey; size:50"`
	Title          string              `json:"title" gorm:"size:100"`
	Slug           string              `json:"slug" gorm:"uniqueIndex; size:100"`
	Thumbnail      string              `json:"thumbnail" gorm:"size:250"`
	Description    string              `json:"description"`
	CategorySerial string              `json:"categorySerial" gorm:"size:50"`
	Category       categories.Category `json:"category" gorm:"foreignKey:CategorySerial; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Images         []ProductImage      `json:"images" gorm:"foreignKey:ProductSerial; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ProductImage struct {
	Serial        string `json:"serial" gorm:"primaryKey; size:50"`
	Url           string `json:"url" gorm:"size:250"`
	ProductSerial string `json:"productSerial" gorm:"size:50"`
}
