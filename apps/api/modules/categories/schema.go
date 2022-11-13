package categories

type Category struct {
	Serial string `json:"serial" gorm:"primaryKey;size:50"`
	Title  string `json:"title" gorm:"size:100"`
	Slug   string `json:"slug" gorm:"size:100"`
	Icon   string `json:"icon" gorm:"size:250"`
}
