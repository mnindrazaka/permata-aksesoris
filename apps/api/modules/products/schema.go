package products

type Product struct {
	Serial      string `json:"serial" gorm:"primaryKey"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}
