package pkg

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"uuid"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lucaastorres7/karakuri/internal/config"
)

type UserClaims struct {
	UserId uuid.UUID `json:"userId"`
	jwt.RegisteredClaims
}

type contextKey string
const userClaimsKey contextKey = "user_claims"

func CreateToken(secret string, userId uuid.UUID) (string, error) {
	expiration := time.Second * time.Duration(config.Env.TokenExpirationSeconds)

	claims := UserClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ExtractBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("user is not logged in")
	}

	fields := strings.Split(authHeader, " ")
	if len(fields) != 2 || fields[0] != "Bearer" {
		return "", fmt.Errorf("invalid auth header")
	}

	return fields[1], nil
}

func ValidateToken(tokenStr, secret string) (*UserClaims, error) {
	claims := &UserClaims{}
	
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func SetUserClaims(ctx context.Context, claims *UserClaims) context.Context {
	return context.WithValue(ctx, userClaimsKey, claims)
}

func GetUserClaims(ctx context.Context) (*UserClaims, bool) {
	claim, ok := ctx.Value(userClaimsKey).(*UserClaims)
	return claim, ok
}