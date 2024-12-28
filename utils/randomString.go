package utils

import (
	"crypto/rand"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLen := len(charset)
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)
	for i := 0; i < length; i++ {
		randomBytes[i] = charset[int(randomBytes[i])%charsetLen]
	}
	return string(randomBytes)
}
