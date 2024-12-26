package roulette

import (
	"fmt"
	"sort"
)

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
