package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
),
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   int64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(userID int64, username string, secret []byte) (string, error) {
	claims := Claims{
		UserID: userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
