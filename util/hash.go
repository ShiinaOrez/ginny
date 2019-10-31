package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func PasswordHash(password string) (string, string, error) {
	h := sha256.New()
	h.Write([]byte(password))
	salt, err := generateSalt(6)
	if err != nil {
		return "", "", err
	}
	h.Write(salt)
	passwordHash := hex.EncodeToString(h.Sum(nil))
	return passwordHash, hex.EncodeToString(salt), nil
}

func CheckPassword(password, passwordHash, salt string) bool {
	h := sha256.New()
	h.Write([]byte(password))
	saltBytes, err := hex.DecodeString(salt)
	if err != nil {
		return false
	}
	h.Write(saltBytes)
	return hex.EncodeToString(h.Sum(nil)) == passwordHash
}

func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return []byte{}, err
	}
	return salt, nil
}
