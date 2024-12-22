package simulation

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	// Test case with a simple sequence of numbers
	numbers := []int{0, 32, 15, 19, 4, 21, 2, 25}
	bet := 10.0
	protection := 3

	// Since Run() only prints output, we're mainly testing that it doesn't panic
	Run(numbers, bet, protection)
}

func TestRunMartingale(t *testing.T) {
	// Test case with a simple sequence of numbers
	numbers := []int{0, 32, 15, 19, 4, 21, 2, 25}
	bet := 10.0
	protection := 3

	// Since RunMartingale() only prints output, we're mainly testing that it doesn't panic
	RunMartingale(numbers, bet, protection)
}

func TestPlay(t *testing.T) {
	testCases := []struct {
		name       string
		numbers    []int
		protection int
		expected   []bool
	}{
		{
			name:       "All win",
			numbers:    []int{8, 17, 14, 23},
			protection: 0,
			expected:   []bool{true, true, true},
		},
		{
			name:       "Ordering",
			numbers:    []int{23, 14, 23},
			protection: 0,
			expected:   []bool{true, false},
		},
		{
			name:       "All lose",
			numbers:    []int{35, 1, 18, 34, 1},
			protection: 0,
			expected:   []bool{false, false, false, false},
		},
		{
			name:       "With protection 1",
			numbers:    []int{29, 23, 18, 34, 1},
			protection: 1,
			expected: []bool{
				false, // bet for 1
				true,  // repeat 1
				false, // bet for 18
				false, // repeat 18 (would win if using 29)
			},
		},
		{
			name:       "Protection 5 (winning last)",
			numbers:    []int{3, 29, 34, 34, 34, 34, 34, 1},
			protection: 5,
			expected: []bool{
				false, // bet for 1
				false, // repeat 1 (would win if using 34)
				false, // repeat 1 (would win if using 34)
				false, // repeat 1 (would win if using 34)
				false, // repeat 1 (would win if using 34)
				true,  // repeat 1 (would lose if using 34)
				true,  // bet for 29 (would lose if using 34 or 1)
			},
		},
		{
			name:       "Protection 5",
			numbers:    []int{34, 34, 34, 34, 34, 34, 34, 1},
			protection: 5,
			expected: []bool{
				false, // bet for 1
				false, // repeat 1 (would win if using 34)
				false, // repeat 1 (would win if using 34)
				false, // repeat 1 (would win if using 34)
				false, // repeat 1 (would win if using 34)
				false, // repeat 1 (would win if using 34)
				true,  // bet for 34
			},
		},
		{
			name:       "Single number",
			numbers:    []int{0},
			protection: 3,
			expected:   []bool{},
		},
		{
			name:       "No protection",
			numbers:    []int{22, 28, 17, 7},
			protection: 0,
			expected: []bool{
				false, // bet for 7
				false, // bet for 17 (would win if using 7)
				false, // bet for 28 (would win if using 17)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			results := play(tc.numbers, tc.protection)
			if !reflect.DeepEqual(results, tc.expected) {
				t.Errorf("play() = %v, want %v", results, tc.expected)
			}
		})
	}
}

func TestCompareWithPrevious(t *testing.T) {
	testCases := []struct {
		name     string
		drawn    int
		game     []int
		expected bool
	}{
		{
			name:     "Exact match",
			drawn:    32,
			game:     []int{32},
			expected: true,
		},
		{
			name:     "No match",
			drawn:    0,
			game:     []int{32},
			expected: false,
		},
		{
			name:     "Empty game",
			drawn:    32,
			game:     []int{},
			expected: false,
		},
		{
			name:     "Match with multiple numbers in game",
			drawn:    32,
			game:     []int{0, 22, 32, 15},
			expected: true,
		},
		{
			name:     "No match with multiple numbers in game",
			drawn:    33,
			game:     []int{34, 14, 32, 10},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := checkWin(tc.drawn, tc.game)
			if result != tc.expected {
				t.Errorf("compareWithPrevious(%d, %v) = %v, want %v",
					tc.drawn, tc.game, result, tc.expected)
			}
		})
	}
}
