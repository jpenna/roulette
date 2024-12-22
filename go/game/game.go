package game

import (
	"fmt"

	"elem.com/roulette/roulette"
)

func GetAllBets(drawn int) ([]int, error) {
	targets, err := roulette.GetExpectedFor(drawn)
	if err != nil {
		return nil, fmt.Errorf("error getting expected for %d: %w", drawn, err)
	}

	var bets []int
	for _, target := range targets {
		index := roulette.RouletteNumberToIndex[target]

		var prev, next int

		prev = roulette.RouletteNumbers[func() int {
			if index == 0 {
				return len(roulette.RouletteNumbers) - 1
			}
			return index - 1
		}()]

		next = roulette.RouletteNumbers[func() int {
			if index == len(roulette.RouletteNumbers)-1 {
				return 0
			}
			return index + 1
		}()]

		bets = append(bets, prev, target, next)
	}

	return bets, nil
}
