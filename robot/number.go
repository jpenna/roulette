package robot

import (
	"fmt"

	"elem.com/roulette/game"
)

func MatchNumbers() {
	window := Window{
		TopLeft:     [2]int{3863, 390},
		BottomRight: [2]int{5997, 1591},
	}
	// window.CaptureSize()
	window.SetNumberAreas()

	numberArea, winArea := game.NewDrawnAreas(window.NumberArea, window.WinArea)

	numberCh := make(chan int)

	// Start a single goroutine that continuously reads numbers
	go func() {
		for {
			game.ReadNumber(numberCh, numberArea, winArea)
		}
	}()

	for number := range numberCh {
		fmt.Println("Number:", number)
	}
}
