package robot

import (
	"fmt"

	"elem.com/roulette/game"
)

func MatchNumbers() {
	window := Window{}
	window.CaptureSize()
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
