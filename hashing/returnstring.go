package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncriptPassword(password, hashing string)bool{
	hash := sha256.Sum256([]byte(password))
	hashStr := hex.EncodeToString(hash[:])

	return hashStr == hashing
}