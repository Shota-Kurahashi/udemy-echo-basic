package utils

import (
	"ge-rest-api/src/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.GetEnv("SECRET")))
}
