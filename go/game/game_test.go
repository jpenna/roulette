package game

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetAllBets(t *testing.T) {
	tests := []struct {
		name     string
		drawn    int
		expected []int
		wantErr  bool
	}{
		{
			name:  "Valid number 0",
			drawn: 0,
			expected: []int{34, 14, 32, 10,
				6, 17, 20, 31, 15, 0, 5, 23},
		},
		{
			name:  "Valid number 18 (next to start)",
			drawn: 18,
			expected: []int{5, 6, 22, 19,
				10, 24, 27, 34, 9, 18, 4, 15},
		},
		{
			name:  "Valid number 20 (next to start)",
			drawn: 20,
			expected: []int{2, 20, 10, 6,
				25, 21, 1, 14, 23, 5, 27, 34},
		},
		{
			name:  "Valid number 35 (next to end)",
			drawn: 35,
			expected: []int{15, 12, 8, 9,
				19, 32, 28, 35, 23, 30, 31, 22},
		},
		{
			name:  "Valid number 36",
			drawn: 36,
			expected: []int{16, 36, 1, 12,
				24, 33, 11, 13, 33, 20, 28, 35},
		},
		{
			name:     "Invalid number (negative)",
			drawn:    -1,
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Invalid number (too large)",
			drawn:    37,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllBets(tt.drawn)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("GetAllBets(%v) returned error: %v", tt.drawn, err)
				}
				return
			}

			if len(got) != len(tt.expected) {
				t.Errorf("GetAllBets(%v) of different length = %v, want %v", tt.drawn, got, tt.expected)
			}

			sort.Ints(got)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("GetAllBets(%v) = %v, want %v", tt.drawn, got, tt.expected)
			}
		})
	}
}
