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

// TestGeneratePrivateKeyStr tests the GeneratePrivateKeyStr function.
func TestGeneratePrivateKeyStr(t *testing.T) {
	// Generate a new private key string
	privateKeyStr, err := cryptoutils.GeneratePrivateKeyStr()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to generate private key string: %v", err)
	}

	// Check if the private key string is empty
	if privateKeyStr == "" {
		t.Fatal("Generated private key string is empty")
	}

	// Check if the private key string is a valid hex string
	if _, err := hex.DecodeString(privateKeyStr); err != nil {
		t.Fatalf("Generated private key string is not a valid hex string: %v", err)
	}
}
