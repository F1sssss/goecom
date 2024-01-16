package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// ComparePasswords compares the hashed password with the password entered by the user
func ComparePasswords(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
