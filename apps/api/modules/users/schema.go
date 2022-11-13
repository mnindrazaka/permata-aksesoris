package users

type User struct {
	Serial   string `json:"serial" gorm:"primaryKey; size:50"`
	Email    string `json:"email" gorm:"uniqueIndex; size:50"`
	Password string `json:"password" gorm:"size:50"`
}
