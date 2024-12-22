package game

import (
	"slices"
	"sort"
	"testing"
)

func TestGetAllBets(t *testing.T) {
	tests := []struct {
		name    string
		drawn   int
		targets []int
		all     []int
		wantErr bool
	}{
		{
			name:    "Valid number 0",
			drawn:   0,
			targets: []int{34, 14, 32, 10},
			all: []int{34, 14, 32, 10,
				6, 17, 20, 31, 15, 0, 5, 23},
		},
		{
			name:    "Valid number 18 (next to start)",
			drawn:   18,
			targets: []int{5, 6, 22, 19},
			all: []int{5, 6, 22, 19,
				10, 24, 27, 34, 9, 18, 4, 15},
		},
		{
			name:    "Valid number 20 (next to start)",
			drawn:   20,
			targets: []int{2, 20, 10, 6},
			all: []int{2, 20, 10, 6,
				25, 21, 1, 14, 23, 5, 27, 34},
		},
		{
			name:    "Valid number 35 (next to end)",
			drawn:   35,
			targets: []int{15, 12, 8, 9},
			all: []int{15, 12, 8, 9,
				19, 32, 28, 35, 23, 30, 31, 22},
		},
		{
			name:    "Valid number 36",
			drawn:   36,
			targets: []int{16, 36, 1, 12},
			all: []int{16, 36, 1, 12,
				24, 33, 11, 13, 33, 20, 28, 35},
		},
		{
			name:    "Invalid number (negative)",
			drawn:   -1,
			all:     nil,
			wantErr: true,
		},
		{
			name:    "Invalid number (too large)",
			drawn:   37,
			all:     nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tgs, all, err := GetAllBets(tt.drawn)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("GetAllBets(%v) returned error: %v", tt.drawn, err)
				}
				return
			}

			sort.Ints(tgs)
			sort.Ints(tt.targets)

			if !slices.Equal(tgs, tt.targets) {
				t.Errorf("Targets for %v = %v, want %v", tt.drawn, tgs, tt.targets)
			}

			if len(all) != len(tt.all) {
				t.Errorf("GetAllBets(%v) of different length = %v, want %v", tt.drawn, all, tt.all)
			}

			sort.Ints(all)
			sort.Ints(tt.all)

			if !slices.Equal(all, tt.all) {
				t.Errorf("GetAllBets(%v) = %v, want %v", tt.drawn, all, tt.all)
			}
		})
	}
}
