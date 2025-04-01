package cryptoutils

import (
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

			if tt.valid {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if address == "" {
					t.Errorf("expected a valid address, got empty string")
				}
			} else {
				if err == nil {
					t.Errorf("expected an error, got none")
				}
			}
		})
	}
}
