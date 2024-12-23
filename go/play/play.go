package play

import (
	"fmt"
	"math/rand/v2"
	"time"

	"elem.com/roulette/halt"
	"elem.com/roulette/robot"
)

func Play() {
	window := robot.Window{}
	window.CaptureSize()
	window.CaptureTerminal()

	rouletteMap, err := robot.UseRouletteMap("roulette.json", &window)
	if err != nil {
		fmt.Println("Error loading roulette map:", err)
		return
	}

	maxProtection := requestProtection()
	game := NewGameState(maxProtection)

	ch := make(chan []int)
	go game.RunRobot(ch, &window)

	for targets := range ch {
		halt.IsHalted.Store(false)

		go selectTargets(targets, &window, rouletteMap)

		halt.ListenForHalt()
	}
}

func requestProtection() int {
	fmt.Print("Máximo de proteção (Enter para usar 2): ")
	var maxProtection int = 2
	var input string
	fmt.Scanln(&input)
	if input != "" {
		if n, err := fmt.Sscanf(input, "%d", &maxProtection); err != nil || n != 1 {
			fmt.Println("Valor inválido, usando proteção padrão (2)")
			maxProtection = 2
		}
	}
	return maxProtection
}

func selectTargets(targets []int, window *robot.Window, rouletteMap *robot.RouletteMap) {
	for _, target := range targets {
		if halt.IsHalted.Load() {
			break
		}

		rouletteMap.ClickNumber(target)

		delay := time.Duration(300+rand.Float64()*1000) * time.Millisecond
		time.Sleep(delay)
	}

	if !halt.IsHalted.Load() {
		window.ClickTerminal()
	}

	halt.StopListenForHalt()
}
