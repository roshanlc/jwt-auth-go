// this package contains functions related to password hash generation
// and password verification, and similar functions
package utilities

import "golang.org/x/crypto/bcrypt"

// GenerateHashPassword returns hash of a password
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash returns boolean if provided plain password
// and password hash match
func CheckPasswordHash(plainPassword, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(plainPassword))
	return err == nil
}
