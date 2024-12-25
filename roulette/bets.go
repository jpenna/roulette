package roulette

import (
	"fmt"
	"sort"
)

// var betsMap = map[int][]int{
// 	0:  {34, 14, 32, 10},
// 	1:  {36, 1, 2, 29},
// 	2:  {20, 5, 2, 22},
// 	3:  {35, 4, 33, 6},
// 	4:  {12, 22, 2, 24},
// 	5:  {18, 6, 24, 2},
// 	6:  {5, 20, 12, 17},
// 	7:  {16, 14, 28, 4},
// 	8:  {11, 28, 3, 31}, // Duplication: 3 instead of 35 (12 repeated)
// 	9:  {31, 11, 6, 3},
// 	10: {23, 20, 28, 19},
// 	11: {8, 29, 31, 13},
// 	12: {21, 32, 36, 3},
// 	13: {31, 11, 33, 15},
// 	14: {34, 14, 30, 5},
// 	15: {35, 20, 17, 24},
// 	16: {36, 16, 19, 7},
// 	17: {17, 22, 16, 8},
// 	18: {5, 6, 22, 19},
// 	19: {16, 28, 21, 36},
// 	20: {2, 20, 10, 6},
// 	21: {2, 12, 19, 16}, // Duplication: 2 instead of 21 (4 repeated)
// 	22: {2, 34, 32, 18}, // Duplication: 34 instead of 17 (25 repeated)
// 	23: {32, 23, 7, 14},
// 	24: {27, 22, 7, 26},
// 	25: {27, 22, 2, 26},
// 	26: {29, 0, 23, 34},
// 	27: {24, 25, 13, 26},
// 	28: {8, 12, 19, 24},
// 	29: {26, 11, 1, 18},
// 	30: {14, 30, 13, 16}, // Duplication: 13 instead of 36 (11 repeated)
// 	31: {27, 22, 11, 28}, // Duplication: 27 instead of 13 (36 repeated)
// 	32: {23, 12, 22, 15},
// 	33: {36, 31, 3, 1},
// 	34: {0, 14, 34, 36},
// 	35: {15, 12, 8, 9},
// 	36: {24, 36, 1, 12}, // Duplication: 24 instead of 16 (33 repeated)
// }

var betsMap = map[int][]int{
	0:  {1, 2, 3, 4},
	1:  {1, 2, 3, 4},
	2:  {1, 2, 3, 4},
	3:  {1, 2, 3, 4},
	4:  {1, 2, 3, 4},
	5:  {1, 2, 3, 4},
	6:  {1, 2, 3, 4},
	7:  {1, 2, 3, 4},
	8:  {1, 2, 3, 4},
	9:  {1, 2, 3, 4},
	10: {1, 2, 3, 4},
	11: {1, 2, 3, 4},
	12: {1, 2, 3, 4},
	13: {1, 2, 3, 4},
	14: {1, 2, 3, 4},
	15: {1, 2, 3, 4},
	16: {1, 2, 3, 4},
	17: {1, 2, 3, 4},
	18: {1, 2, 3, 4},
	19: {1, 2, 3, 4},
	20: {1, 2, 3, 4},
	21: {1, 2, 3, 4},
	22: {1, 2, 3, 4},
	23: {1, 2, 3, 4},
	24: {1, 2, 3, 4},
	25: {1, 2, 3, 4},
	26: {1, 2, 3, 4},
	27: {1, 2, 3, 4},
	28: {1, 2, 3, 4},
	29: {1, 2, 3, 4},
	30: {1, 2, 3, 4},
	31: {1, 2, 3, 4},
	32: {1, 2, 3, 4},
	33: {1, 2, 3, 4},
	34: {1, 2, 3, 4},
	35: {1, 2, 3, 4},
	36: {1, 2, 3, 4},
}

func GetTargetBetsFor(number int) ([]int, error) {
	if number < 0 || number > 36 {
		return nil, fmt.Errorf("number out of range: %d", number)
	}

	expected, exists := betsMap[number]
	if !exists {
		return nil, fmt.Errorf("no expectations found for number: %d", number)
	}

	return expected, nil
}

func GetAllBetsFor(drawn int) (targets []int, all []int, err error) {
	targets, err = GetTargetBetsFor(drawn)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting expected for %d: %w", drawn, err)
	}

	for _, target := range targets {
		index := RouletteNumberToIndex[target]

		var prev, next int

		prev = RouletteNumbers[func() int {
			if index == 0 {
				return len(RouletteNumbers) - 1
			}
			return index - 1
		}()]

		next = RouletteNumbers[func() int {
			if index == len(RouletteNumbers)-1 {
				return 0
			}
			return index + 1
		}()]

		all = append(all, prev, target, next)
	}

	return targets, all, nil
}

func FindDuplicatedBets() {
	for key := range betsMap {
		targets, all, err := GetAllBetsFor(key)
		if err != nil {
			fmt.Printf("Error processing number %d: %v\n", key, err)
			continue
		}

		sort.Ints(all)

		prev := -1
		for _, num := range all {
			// log.Printf("Processing number %d\n", num)
			if prev == num {
				fmt.Printf("%d: Number %d is repeated\n (target: %+v)\n", key, num, targets)
			}
			prev = num
		}
	}
}
