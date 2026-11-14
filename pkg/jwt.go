package pkg

import (
	"time"
	"uuid"

	"github.com/golang-jwt/jwt/v5"
)


func CreateToken(secret string, userId uuid.UUID) (string, error) {
	expiration := time.Second

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}