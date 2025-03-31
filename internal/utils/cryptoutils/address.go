package cryptoutils

import (
	"encoding/base32"
	"encoding/hex"
	"strings"

	"lukechampine.com/blake3"
)

// GenerateWalsletAddress generates a wallet address from the given public key.
func GenerateWalletAddress(publicKey string) (string, error) {
	// Convert the public key from hex to bytes
	publicKeyBytes, err := hex.DecodeString(publicKey)

	// Check for errors
	if err != nil {
		return "", err
	}

	// In my own implementation,
	// I would use Blake3 for deriving the public key to generate the wallet address
	hash := blake3.Sum256(publicKeyBytes)
	// Get the base32 encoding of the hash
	address := base32.StdEncoding.EncodeToString(hash[:])

	// Remove the padding characters from the base32 string
	return strings.ReplaceAll(address, "=", ""), nil
}
