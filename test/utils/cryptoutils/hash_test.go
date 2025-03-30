package cryptoutils

import (
	"mavyfaby/blockchain/internal/utils/cryptoutils"
	"testing"
)

// TestHashSHA256 tests the HashSHA256 function
func TestHashSHA256(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "plain string",
			input:    "This is a plain string",
			expected: "7fcfeaf04c73596c51a257a2cc9198b4f2de968c1f06eec2ad4713d335f43d1d",
		},
		{
			name:     "special characters",
			input:    "Hello, Mavyfaby! @2025",
			expected: "0dd00af5702882c9393e941b7c3ddd41d5511e29f69711994c60c2d85a6c395f",
		},
		{
			name:     "whitespaces",
			input:    "  Hello, Mavyfaby!  ",
			expected: "9fc678bc49d7e19ed1755dde5ca1c4aeab6064cd3c51af707c12f0fbff454676",
		},
	}

	// Iterate over each test case
	for _, test := range tests {
		// Run the test case
		t.Run(test.name, func(t *testing.T) {
			// Call the function to test
			result := cryptoutils.HashSHA256(test.input)

			// Check if the result matches the expected output
			if result != test.expected {
				// If not, fail the test and print the expected and actual values
				t.Errorf("expected %s, got %s", test.expected, result)
			}
		})
	}
}
