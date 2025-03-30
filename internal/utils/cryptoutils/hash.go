package cryptoutils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashSHA256 takes a string input and returns its SHA256 hash in hexadecimal format
func HashSHA256(input string) string {
	// Create a new SHA256 hash instance
	h := sha256.New()
	// Write the input string to the hash
	h.Write([]byte(input))

	// Calculate the hash and convert it to a byte slice
	return hex.EncodeToString(h.Sum(nil))
}
