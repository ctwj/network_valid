package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStr(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"length 10", 10},
		{"length 20", 20},
		{"length 0", 0},
		{"length 1", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandStr(tt.length)
			assert.Equal(t, tt.length, len(result), "length should match")

			// Check characters are from the expected set
			validChars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
			for _, c := range result {
				assert.True(t, strings.ContainsRune(validChars, c), "character should be from valid set")
			}
		})
	}

	// Test uniqueness
	t.Run("uniqueness", func(t *testing.T) {
		results := make(map[string]bool)
		for i := 0; i < 100; i++ {
			s := RandStr(16)
			assert.False(t, results[s], "should generate unique strings")
			results[s] = true
		}
	})
}

func TestRandUpperStr(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"length 10", 10},
		{"length 20", 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandUpperStr(tt.length)
			assert.Equal(t, tt.length, len(result))

			// Check characters are uppercase or digits
			validChars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
			for _, c := range result {
				assert.True(t, strings.ContainsRune(validChars, c))
			}
		})
	}
}

func TestRandLowerStr(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"length 10", 10},
		{"length 20", 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandLowerStr(tt.length)
			assert.Equal(t, tt.length, len(result))

			// Check characters are lowercase or digits
			validChars := "0123456789abcdefghijklmnopqrstuvwxyz"
			for _, c := range result {
				assert.True(t, strings.ContainsRune(validChars, c))
			}
		})
	}
}

func TestRandNumStr(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"length 10", 10},
		{"length 20", 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandNumStr(tt.length)
			assert.Equal(t, tt.length, len(result))

			// Check all characters are digits
			validChars := "0123456789"
			for _, c := range result {
				assert.True(t, strings.ContainsRune(validChars, c))
			}
		})
	}
}

func TestIn(t *testing.T) {
	tests := []struct {
		name     string
		target   string
		array    []string
		expected bool
	}{
		{"found", "apple", []string{"apple", "banana", "cherry"}, true},
		{"not found", "grape", []string{"apple", "banana", "cherry"}, false},
		{"empty array", "apple", []string{}, false},
		{"single element found", "apple", []string{"apple"}, true},
		{"single element not found", "banana", []string{"apple"}, false},
		{"case sensitive", "Apple", []string{"apple", "banana"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := In(tt.target, tt.array)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetRsaKey(t *testing.T) {
	status, publicKey, privateKey := GetRsaKey()

	assert.True(t, status, "should generate RSA key successfully")
	assert.NotEmpty(t, publicKey, "public key should not be empty")
	assert.NotEmpty(t, privateKey, "private key should not be empty")

	assert.True(t, strings.Contains(publicKey, "BEGIN PUBLIC KEY"))
	assert.True(t, strings.Contains(publicKey, "END PUBLIC KEY"))
	assert.True(t, strings.Contains(privateKey, "BEGIN RSA PRIVATE KEY"))
	assert.True(t, strings.Contains(privateKey, "END RSA PRIVATE KEY"))
}
