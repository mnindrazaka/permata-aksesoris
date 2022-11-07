package middlewares

import (
	"fmt"
	"net/http"
	"permata-aksesoris/apps/api/utils"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func NewAuthenticateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenWithBearer := r.Header.Get("authorization")

		if tokenWithBearer == "" {
			utils.WriteUnauthorizedResponse(w, fmt.Errorf("no authorization header found"))
			return
		}

		splittedToken := strings.Split(tokenWithBearer, " ")

		if len(splittedToken) != 2 {
			utils.WriteUnauthorizedResponse(w, fmt.Errorf("authorization header format is not correct"))
			return
		}

		tokenString := splittedToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return []byte("SECRET"), nil
		})

		if err != nil {
			utils.WriteUnauthorizedResponse(w, err)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			next(w, r)
		} else {
			utils.WriteUnauthorizedResponse(w, err)
		}
	}
}
