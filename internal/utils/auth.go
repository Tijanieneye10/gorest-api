package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId int64, username string, secret []byte) (string, error) {
	
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := claims.SignedString(secret)

	if err != nil {
		return "", err
	}

	return string(token), nil
}
