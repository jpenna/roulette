package robot

import (
	"fmt"

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

	numberCh := make(chan int)
	go numberArea.ReadNumber(numberCh)

	for number := range numberCh {
		fmt.Println("Number:", number)
	}
}
