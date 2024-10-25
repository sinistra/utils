package utils

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/go-passwd/validator"
	"golang.org/x/crypto/bcrypt"
)

// IsValidPassword returns true if the password passes the validation tests.
func IsValidPassword(password string) bool {
	passwordValidator := validator.New(
		validator.MinLength(8, nil),
		validator.ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 1, nil),
		validator.ContainsAtLeast("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1, nil),
		validator.ContainsAtLeast("0123456789", 1, nil),
		validator.CommonPassword(nil),
	)

	err := passwordValidator.Validate(password)

	return err == nil
}

// GeneratePasswordHash returns a hash of the password or an error.
func GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// ComparePasswordHash returns true if the password compares to the hash.
func ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateSecureToken creates a string of random characters.
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	return hex.EncodeToString(b)
}
