package utils

import (
	"crypto/rand"
	"math/big"
)

const (
	charList string = "abcdefghijklmnopqrstuvwxyzABCDEFHFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func RandomString(size int) string {
	stringLength := min(size, len(charList))
	theKey := make([]byte, stringLength)
	for key := 0; key < stringLength; key++ {
		res, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charList))))
		keyGen := charList[res.Int64()]
		theKey[key] = keyGen
	}
	return string(theKey)
}
