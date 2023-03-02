package utilities

import "golang.org/x/crypto/bcrypt"

// GenerateHashPassword returns hash of a password
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
