package users

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
)

type usecase struct {
	repository Repository
}

type Usecase interface {
	login(email string, password string) (string, error)
}

func NewUsecase(repository Repository) Usecase {
	return usecase{repository}
}

func (usecase usecase) login(email string, password string) (string, error) {
	user, err := usecase.repository.getUserByEmailAndPassword(email, password)

	if err != nil {
		return "", err
	}

	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	var claims jwt.MapClaims
	if err := json.Unmarshal(data, &claims); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("SECRET"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
