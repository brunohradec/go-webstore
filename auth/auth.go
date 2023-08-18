package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
