package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Generate hashed password
func GeneratePassword(password string) ([]byte, error) {
	// 10 is the cost representing the password is hashed 2^10 times i.e the hash is also hashed multiple times
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatalf("Error generating password hash: %v", err)
		return nil, err
	}

	return hash, nil
}

// compare password string with hash
func ComparePassword(hash []byte, password string) error {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return err
	}

	return nil
}
