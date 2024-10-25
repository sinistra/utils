package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{
			name:     "Common password",
			password: "Passw0rd",
			want:     false,
		},
		{
			name:     "no lower case",
			password: "MYLONG1STPASSWORD",
			want:     false,
		},
		{
			name:     "no upper case",
			password: "mylong1stpassword",
			want:     false,
		},
		{
			name:     "no digits",
			password: "mylongPASSWORD",
			want:     false,
		},
		{
			name:     "valid password",
			password: "mylong3stPASSWORD",
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidPassword(tt.password); got != tt.want {
				t.Errorf("IsValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPasswordHash(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
		wantErr  bool
	}{
		{
			name:     "success",
			password: "password",
			want:     true,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := GeneratePasswordHash(tt.password)
			assert.NoError(t, err)

			got := ComparePasswordHash(tt.password, hash)
			if got != tt.want {
				t.Errorf("GeneratePasswordHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateSecureToken(t *testing.T) {
	tests := []struct {
		name   string
		length int
		want   int
	}{
		{
			name:   "empty",
			length: 0,
			want:   0,
		},
		{
			name:   "one",
			length: 1,
			want:   2,
		},
		{
			name:   "two",
			length: 2,
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateSecureToken(tt.length)
			if len(got) != tt.length*2 {
				t.Errorf("GenerateSecureToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
