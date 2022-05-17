package utils

import (
	"reflect"
	"testing"
)

func Test_max(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want int
	}{
		{
			name: "TestMax#1",
			s:    []int{4, 5, 6},
			want: 6,
		},
		{
			name: "TestMax#0",
			s:    []int{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want int
	}{
		{
			name: "TestMin#1",
			s:    []int{5, 4, 5, 6},
			want: 4,
		},
		{
			name: "TestMin#0",
			s:    []int{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
