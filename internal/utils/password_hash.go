package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hasPassword), nil
}

func CheckPasswordHash(storedPassword, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(storedPassword))
	return err == nil

}
