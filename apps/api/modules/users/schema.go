package users

type User struct {
	Serial   string `json:"serial" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
