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
			name:  "ContainsStringYes",
			items: []string{"foo", "bar"},
			item:  "bar",
			want:  true,
		},
		{
			name:  "ContainsStringNo",
			items: []string{"foo", "bar"},
			item:  "oh",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.items, tt.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt(t *testing.T) {
	tests := []struct {
		name  string
		items []int
		item  int
		want  bool
	}{
		{
			name:  "ContainsIntYes",
			items: []int{1, 2, 3},
			item:  1,
			want:  true,
		},
		{
			name:  "ContainsIntNo",
			items: []int{1, 2, 3},
			item:  4,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.items, tt.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
