package robot

import (
	"errors"
	"fmt"
	"time"

	"elem.com/roulette/game"
)

func MatchNumbers() {

	window := Window{
		// TopLeft:     [2]int{4151, 348},
		// BottomRight: [2]int{5800, 1268},
	}
	window.CaptureSize()
	window.SetNumberArea()

	numberArea := game.NewNumberArea(window.NumberArea)

	for {
		time.Sleep(1000 * time.Millisecond)

		number, err := numberArea.CaptureNumber()
		if err != nil {
			if errors.Is(err, game.ErrNoNumber) {
				fmt.Println("-")
				continue
			}

			if errors.Is(err, game.ErrWrongColor) {
				fmt.Printf("\033[41m%v\033[0m\n", err)
				continue
			}

			fmt.Printf("\033[41mError capturing number: %v\033[0m\n", err)
			continue
		}

		fmt.Println("Number:", number)
	}
}
