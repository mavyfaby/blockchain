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

// GeneratePublicKey generates the public key from the given private key.
func GeneratePublicKey(privateKey *ecdsa.PrivateKey) string {
	// Generate the public key from the private key
	compressed := elliptic.MarshalCompressed(privateKey.Curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)
	// Return the public key
	return hex.EncodeToString(compressed)
}

// GenerateKeyPair generates a new ECDSA key pair (private and public keys) and returns them as strings.
func GenerateKeyPair() (string, string, error) {
	// Generate a new private key
	privateKey, err := GeneratePrivateKey()

	// Check for errors
	if err != nil {
		return "", "", err
	}

	// Generate the public key from the private key
	publicKey := GeneratePublicKey(privateKey)

	// Return the private and public keys as strings
	return hex.EncodeToString(privateKey.D.Bytes()), publicKey, nil
}
