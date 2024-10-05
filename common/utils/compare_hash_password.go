package utils

import "golang.org/x/crypto/bcrypt"

func CompareHashPassword(hashedPassword string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}

	return true
}
