package robot

import (
	"fmt"
	"math/rand"
	"time"

	"elem.com/roulette/play"
)

func Play() {
	window := Window{}
	window.Capture()
	window.CaptureTerminal()
	rouletteMap, err := UseRouletteMap("roulette.json", &window)
	if err != nil {
		fmt.Println("Error loading roulette map:", err)
		return
	}

	ch := make(chan []int)

	game := play.NewGameState()

	go game.RunRobot(ch)

	for targets := range ch {
		for _, target := range targets {
			rouletteMap.ClickNumber(target)

			delay := time.Duration(800+rand.Float64()*1000) * time.Millisecond
			time.Sleep(delay)
		}

		window.ClickTerminal()
	}
}
