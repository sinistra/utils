// Package utils is a collection of often used utilities.
package utils

import "net/mail"

// IsValidEmail returns true if the string can be parsed as an email address.
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
