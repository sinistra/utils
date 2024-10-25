package utils

import "github.com/google/uuid"

// IsValidUUID checks that a string can be parsed as a UUID.
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
