package game

import (
	"fmt"

	"elem.com/roulette/roulette"
	"github.com/go-vgo/robotgo"
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

func IsReadyToBet(x, y int) (bool, error) {
	img, err := robotgo.CaptureImg(x, y, 1, 1)
	if err != nil {
		return false, fmt.Errorf("error capturing image: %w", err)
	}

	pixel := img.At(0, 0)
	r, g, b, _ := pixel.RGBA()

	greenMargin := g - 50 // 50 is the margin of error

	// Check if green component is significantly higher than red and blue
	return greenMargin > r && greenMargin > b, nil
}
