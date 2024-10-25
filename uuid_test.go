package utils

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidUUID(t *testing.T) {
	tests := []struct {
		name string
		u    string
		want bool
	}{
		{
			name: "success",
			u:    uuid.New().String(),
			want: true,
		},
		{
			name: "failure",
			u:    "some-string",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidUUID(tt.u)
			assert.Equalf(t, tt.want, got, "IsValidUUID(%v)", tt.u)
		})
	}
}
