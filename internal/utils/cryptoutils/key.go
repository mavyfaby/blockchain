package cryptoutils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
)

// GeneratePrivateKey generates a new ECDSA private key using the P-256 curve.
func GeneratePrivateKey() (*ecdsa.PrivateKey, error) {
	// Generate a new private key using the ECDSA algorithm
	curve := elliptic.P256()
	// Generate a new private key
	return ecdsa.GenerateKey(curve, rand.Reader)
}

// GeneratePrivateKeyStr generates a new ECDSA private key and returns it as a string.
func GeneratePrivateKeyStr() (string, error) {
	// Generate a new private key
	privateKey, err := GeneratePrivateKey()

	// Check for errors
	if err != nil {
		return "", err
	}

	// Convert the private key to bytes and then to a hex string
	return hex.EncodeToString(privateKey.D.Bytes()), nil
}
