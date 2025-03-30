package cryptoutils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"mavyfaby/blockchain/internal/utils/cryptoutils"
	"testing"
)

// TestGeneratePrivateKey tests the GeneratePrivateKey function.
func TestGeneratePrivateKey(t *testing.T) {
	// Generate a new private key
	privateKey, err := cryptoutils.GeneratePrivateKey()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Check if the private key is nil
	if privateKey == nil {
		t.Fatal("Generated private key is nil")
	}

	// Check if the private key is of the correct type
	if _, ok := privateKey.Public().(*ecdsa.PublicKey); !ok {
		t.Fatal("Generated private key is not of type *ecdsa.PrivateKey")
	}
}

// TestGeneratePublicKey tests the GeneratePublicKey function.
func TestGeneratePublicKey(t *testing.T) {
	// Generate a new private key
	privateKey, err := cryptoutils.GeneratePrivateKey()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Generate the public key from the private key
	publicKey := cryptoutils.GeneratePublicKey(privateKey)

	// Check if the public key is empty
	if publicKey == "" {
		t.Fatal("Generated public key is empty")
	}

	// Check if the public key is of the correct type
	if _, err := hex.DecodeString(publicKey); err != nil {
		t.Fatalf("Generated public key is not of type string: %v", err)
	}
}

// TestGenerateKeyPair tests the GenerateKeyPair function.
func TestGenerateKeyPair(t *testing.T) {
	// Generate a new key pair
	privateKey, publicKey, err := cryptoutils.GenerateKeyPair()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Check if the private key is empty
	if privateKey == "" {
		t.Fatal("Generated private key is empty")
	}

	// Check if the public key is empty
	if publicKey == "" {
		t.Fatal("Generated public key is empty")
	}

	// Check if the private key is of the correct type
	if _, err := hex.DecodeString(privateKey); err != nil {
		t.Fatalf("Generated private key is not of type string: %v", err)
	}

	// Check if the public key is of the correct type
	if _, err := hex.DecodeString(publicKey); err != nil {
		t.Fatalf("Generated public key is not of type string: %v", err)
	}
}
