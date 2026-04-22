package api

import (
	"testing"

	"github.com/astaxie/beego/validation"
	"github.com/stretchr/testify/assert"
)

func TestUnEncrypt_Valid(t *testing.T) {
	tests := []struct {
		name    string
		input   UnEncrypt
		hasErr  bool
		errField string
	}{
		{
			name: "valid input",
			input: UnEncrypt{
				Appkey:    "abc123DEF456",
				Version:   "1.0",
				Sign:      "sign123ABC",
				Action:    "login",
				Timestamp: 1234567890,
				Mac:       "ABC123DEF456",
			},
			hasErr: false,
		},
		{
			name: "invalid appkey with special chars",
			input: UnEncrypt{
				Appkey:    "abc@123",
				Version:   "1.0",
				Sign:      "sign123",
				Timestamp: 1234567890,
				Mac:       "ABC123",
			},
			hasErr:   true,
			errField: "Appkey",
		},
		{
			name: "invalid version",
			input: UnEncrypt{
				Appkey:    "abc123",
				Version:   "abc",
				Sign:      "sign123",
				Timestamp: 1234567890,
				Mac:       "ABC123",
			},
			hasErr:   true,
			errField: "Version",
		},
		{
			name: "invalid sign",
			input: UnEncrypt{
				Appkey:    "abc123",
				Version:   "1.0",
				Sign:      "sign@123",
				Timestamp: 1234567890,
				Mac:       "ABC123",
			},
			hasErr:   true,
			errField: "Sign",
		},
		{
			name: "invalid mac",
			input: UnEncrypt{
				Appkey:    "abc123",
				Version:   "1.0",
				Sign:      "sign123",
				Timestamp: 1234567890,
				Mac:       "MAC-123",
			},
			hasErr:   true,
			errField: "Mac",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validation.Validation{}
			tt.input.Valid(&v)

			if tt.hasErr {
				assert.True(t, v.HasErrors())
				if tt.errField != "" {
					found := false
					for _, err := range v.Errors {
						if err.Field == tt.errField {
							found = true
							break
						}
					}
					assert.True(t, found, "expected error for field %s", tt.errField)
				}
			} else {
				assert.False(t, v.HasErrors())
			}
		})
	}
}

func TestEncrypt_Valid(t *testing.T) {
	tests := []struct {
		name    string
		input   Encrypt
		hasErr  bool
		errField string
	}{
		{
			name: "valid input",
			input: Encrypt{
				Signal:     "signal123",
				Sign:       "sign123ABC",
				Encrypt:    "encrypt123",
				Timestamp:  1234567890,
				Ciphertext: "ciphertext",
			},
			hasErr: false,
		},
		{
			name: "invalid signal",
			input: Encrypt{
				Signal:     "signal@123",
				Sign:       "sign123",
				Encrypt:    "encrypt123",
				Timestamp:  1234567890,
				Ciphertext: "ciphertext",
			},
			hasErr:   true,
			errField: "Signal",
		},
		{
			name: "invalid sign",
			input: Encrypt{
				Signal:     "signal123",
				Sign:       "sign-123",
				Encrypt:    "encrypt123",
				Timestamp:  1234567890,
				Ciphertext: "ciphertext",
			},
			hasErr:   true,
			errField: "Sign",
		},
		{
			name: "invalid encrypt",
			input: Encrypt{
				Signal:     "signal123",
				Sign:       "sign123",
				Encrypt:    "encrypt@123",
				Timestamp:  1234567890,
				Ciphertext: "ciphertext",
			},
			hasErr:   true,
			errField: "Encrypt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validation.Validation{}
			tt.input.Valid(&v)

			if tt.hasErr {
				assert.True(t, v.HasErrors())
			} else {
				assert.False(t, v.HasErrors())
			}
		})
	}
}

func TestRegisterParam_Valid(t *testing.T) {
	tests := []struct {
		name    string
		input   RegisterParam
		hasErr  bool
		errField string
	}{
		{
			name: "valid input",
			input: RegisterParam{
				User:    "testuser123",
				Pwd:     "password123",
				Pwd2:    "password123",
				Email:   "test@example.com",
				Captcha: "ABC123",
			},
			hasErr: false,
		},
		{
			name: "invalid user too short",
			input: RegisterParam{
				User:    "abc",
				Pwd:     "password123",
				Pwd2:    "password123",
				Email:   "test@example.com",
				Captcha: "ABC123",
			},
			hasErr:   true,
			errField: "User",
		},
		{
			name: "invalid user too long",
			input: RegisterParam{
				User:    "abcdefghijklmnopqrstuvwxyz12345678901234567890",
				Pwd:     "password123",
				Pwd2:    "password123",
				Email:   "test@example.com",
				Captcha: "ABC123",
			},
			hasErr:   true,
			errField: "User",
		},
		{
			name: "invalid password",
			input: RegisterParam{
				User:    "testuser123",
				Pwd:     "pwd",
				Pwd2:    "pwd",
				Email:   "test@example.com",
				Captcha: "ABC123",
			},
			hasErr:   true,
			errField: "Pwd",
		},
		{
			name: "invalid email",
			input: RegisterParam{
				User:    "testuser123",
				Pwd:     "password123",
				Pwd2:    "password123",
				Email:   "invalid-email",
				Captcha: "ABC123",
			},
			hasErr:   true,
			errField: "Email",
		},
		{
			name: "invalid captcha",
			input: RegisterParam{
				User:    "testuser123",
				Pwd:     "password123",
				Pwd2:    "password123",
				Email:   "test@example.com",
				Captcha: "CAP-123",
			},
			hasErr:   true,
			errField: "Captcha",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validation.Validation{}
			tt.input.Valid(&v)

			if tt.hasErr {
				assert.True(t, v.HasErrors())
			} else {
				assert.False(t, v.HasErrors())
			}
		})
	}
}
