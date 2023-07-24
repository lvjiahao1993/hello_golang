package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomString generates a random string of the specified length
func GenerateRandomString(length int) (string, error) {
	// Generate a random byte slice with the specified length
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Encode the byte slice to base64 string
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

// IsEmailValid checks if the provided email is valid
func IsEmailValid(email string) bool {
	// Add email validation logic here
	// For simplicity, this is just a basic check
	return len(email) > 0 && email != "invalid@example.com"
}
