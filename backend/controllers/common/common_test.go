package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	token1 := GetToken()
	token2 := GetToken()

	assert.NotEmpty(t, token1, "token should not be empty")
	assert.Len(t, token1, 32, "MD5 token should be 32 characters")
	assert.NotEqual(t, token1, token2, "tokens should be unique")
}

func TestGetStringMd5(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", "d41d8cd98f00b204e9800998ecf8427e"},
		{"hello", "hello", "5d41402abc4b2a76b9719d911017c592"},
		{"password", "password", "5f4dcc3b5aa765d61d8327deb882cf99"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetStringMd5(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Test consistency
	t.Run("consistency", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			assert.Equal(t, GetStringMd5("test"), GetStringMd5("test"))
		}
	})
}

func TestVerifyEmailFormat(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"valid email", "test@example.com", true},
		{"valid with subdomain", "test@sub.example.com", true},
		{"valid with plus", "test+tag@example.com", true},
		{"valid with dots", "test.name@example.com", true},
		{"invalid no @", "testexample.com", false},
		{"invalid no domain", "test@", false},
		{"invalid no local", "@example.com", false},
		{"empty string", "", false},
		{"invalid spaces", "test @example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := VerifyEmailFormat(tt.email)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAesEncryptDecrypt(t *testing.T) {
	key := []byte("0123456789abcdef") // 16 bytes for AES-128
	iv := []byte("abcdef0123456789")

	tests := []struct {
		name string
		text string
	}{
		{"simple text", "hello world"},
		{"empty text", ""},
		{"long text", "this is a longer text that needs to be encrypted and decrypted"},
		{"special chars", "!@#$%^&*()_+-=[]{}|;':\",./<>?"},
		{"chinese", "中文测试"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encrypted, err := AesEncrypt([]byte(tt.text), key, iv)
			assert.NoError(t, err)
			assert.NotEmpty(t, encrypted)

			decrypted, err := AesDecrypt(encrypted, key, iv)
			assert.NoError(t, err)
			assert.Equal(t, tt.text, decrypted)
		})
	}
}

func TestGetInterfaceToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int
	}{
		{"int", 42, 42},
		{"int8", int8(42), 42},
		{"int16", int16(42), 42},
		{"int32", int32(42), 42},
		{"int64", int64(42), 42},
		{"uint", uint(42), 42},
		{"uint8", uint8(42), 42},
		{"uint16", uint16(42), 42},
		{"uint32", uint32(42), 42},
		{"uint64", uint64(42), 42},
		{"float32", float32(42.5), 42},
		{"float64", float64(42.5), 42},
		{"string number", "42", 42},
		{"string float", "42.5", 42},
		{"nil", nil, 0},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetInterfaceToInt(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetFloatLen(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected int
	}{
		{"no decimal", 42.0, 0},
		{"one decimal", 42.5, 1},
		{"two decimals", 42.55, 2},
		{"three decimals", 42.555, 3},
		{"integer", 100.0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetFloatLen(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetVersionString(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"no decimal", 1.0, "1.00"},
		{"one decimal", 1.1, "1.10"},
		{"two decimals", 1.11, "1.11"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetVersionString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetAddrIp(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"with port", "192.168.1.1:8080", "192.168.1.1"},
		{"localhost with port", "127.0.0.1:3000", "127.0.0.1"},
		{"no port", "192.168.1.1", ""},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetAddrIp(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsValueNil(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"nil value", nil, true},
		{"non-nil string", "hello", false},
		{"non-nil int", 42, false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValueNil(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
