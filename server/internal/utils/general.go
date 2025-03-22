package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func GenerateRandomString(length int) (string, error) {
	const possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	result := make([]byte, length)
	for i, b := range bytes {
		result[i] = possible[int(b)%len(possible)]
	}
	return string(result), nil
}

func Sha256Hash(plain string) string {
	hash := sha256.Sum256([]byte(plain))
	return hex.EncodeToString(hash[:])
}

func Base64Encode(input []byte) string {
	encoded := base64.RawURLEncoding.EncodeToString(input)
	return encoded
}
