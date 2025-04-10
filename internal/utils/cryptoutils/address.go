package cryptoutils

import (
	"encoding/base32"
	"encoding/hex"
	"errors"

	"lukechampine.com/blake3"
)

// GenerateWalsletAddress generates a wallet address from the given public key.
func GenerateWalletAddress(publicKey string) (string, error) {
	// Check if publicKey is empty
	if publicKey == "" {
		return "", errors.New("public key cannot be empty")
	}

	// Check if publicKey is a valid hex string
	if _, err := hex.DecodeString(publicKey); err != nil {
		return "", errors.New("invalid public key format")
	}

	// Convert the public key from hex to bytes
	publicKeyBytes, err := hex.DecodeString(publicKey)

	// Check for errors
	if err != nil {
		return "", err
	}

	// In my own implementation,
	// I would use Blake3 for deriving the public key to generate the wallet address
	hash := blake3.Sum256(publicKeyBytes)
	// Encode the hash to base32 without padding
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hash[:]), nil
}
