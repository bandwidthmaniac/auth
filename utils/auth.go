package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Generates a cryptographically secure password
// from a given plain text password.
//
// https://stackoverflow.com/a/23259804/7469926
func HashPassword(password []byte) (string, error) {
	hashedPassowrd, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassowrd), nil
}

// Compares provided `hashedPassword` & `password` for truthiness.
//
// https://stackoverflow.com/a/23259804/7469926
func ComparePassword(hashedPassword []byte, password []byte) error {
	compareError := bcrypt.CompareHashAndPassword(hashedPassword, password)

	if compareError != nil {
		return compareError
	}

	return nil
}
