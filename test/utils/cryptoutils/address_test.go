package cryptoutils

import (
	"encoding/base32"
	"mavyfaby/blockchain/internal/utils/cryptoutils"
	"testing"
)

func TestGenerateWalletAddress(t *testing.T) {
	tests := []struct {
		name  string
		input string
		valid bool
	}{
		{
			name:  "Should not empty",
			input: "",
			valid: false,
		},
		{
			name:  "Valid Public Key",
			input: "0249ff062bb955759391dbfec8344035e6a55f2dc9402631dbcaf37c1bea9eaccb",
			valid: true,
		},
		{
			name:  "Invalid Public Key",
			input: "invalidpublickey",
			valid: false,
		},
		{
			name:  "Invalid Public Key Length",
			input: "344035e6a55f2dc940263",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			address, err := cryptoutils.GenerateWalletAddress(tt.input)

			// Check if the address is not empty
			if tt.valid {
				// Check if the address is not empty
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}

				// Check if the address is not empty
				if address == "" {
					t.Errorf("expected a valid address, got empty string")
				}

				// Check if the address is a valid base32 string
				decoded, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(address)

				// if has error then it is not a valid base32 string
				if err != nil {
					t.Errorf("expected a valid base32 string, got error %v", err)
				} else if len(decoded) == 0 { // Check if the decoded result is not empty
					t.Errorf("expected a valid base32 string, got empty decoded result")
				}
			} else {
				if err == nil {
					t.Errorf("expected an error, got none")
				}
			}
		})
	}
}
