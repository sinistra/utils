package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{
			name:  "success",
			email: "someone@example.com",
			want:  true,
		},
		{
			name:  "fail",
			email: "someone",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsValidEmail(tt.email), "IsValidEmail(%v)", tt.email)
		})
	}
}
