package users

import "gorm.io/gorm"

type repository struct {
	con *gorm.DB
}

type Repository interface {
	getUserByEmailAndPassword(email string, password string) (User, error)
}

func NewRepository(con *gorm.DB) Repository {
	return repository{con}
}

func (repository repository) getUserByEmailAndPassword(email string, password string) (User, error) {
	var user User
	result := repository.con.Model(User{}).Where(User{Email: email, Password: password}).First(&user)
	return user, result.Error
}
