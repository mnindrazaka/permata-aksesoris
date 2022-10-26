package categories

type Category struct {
	Serial string `json:"serial" gorm:"primaryKey"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
	Icon   string `json:"icon"`
}
