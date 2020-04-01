package security

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)


func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword verify the hashed password
func VerifyPassword(hashedPassword, password string) error {
	fmt.Println("Passwordxs", password, hashedPassword)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}