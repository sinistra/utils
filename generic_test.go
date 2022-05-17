package utils

import (
	"testing"
)

func TestContainsString(t *testing.T) {
	tests := []struct {
		name  string
		items []string
		item  string
		want  bool
	}{
		{
			name:  "ContainsString",
			items: []string{"foo", "bar"},
			item:  "bar",
			want:  true,
		},
		{
			name:  "ContainsString",
			items: []string{"foo", "bar"},
			item:  "oh",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains[string](tt.items, tt.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
