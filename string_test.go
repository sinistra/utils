package utils

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "TestRandomString32",
			size: 32,
		},
		{
			name: "TestRandomString64",
			size: 64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.size)
			if len(got) != tt.size {
				t.Errorf("RandomString() = %v, want %v", got, tt.size)
			}
		})
	}
}
