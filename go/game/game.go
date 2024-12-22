package game

import (
	"fmt"

	"elem.com/roulette/roulette"
)

func GetAllBets(drawn int) (targets []int, all []int, err error) {
	targets, err = roulette.GetTargetBetsFor(drawn)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting expected for %d: %w", drawn, err)
	}

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

		all = append(all, prev, target, next)
	}

	return targets, all, nil
}
