package pkg

import (
	"time"
	"uuid"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lucaastorres7/karakuri/internal/config"
)


func CreateToken(secret string, userId uuid.UUID) (string, error) {
	expiration := time.Second * time.Duration(config.Env.TokenExpirationSeconds)

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