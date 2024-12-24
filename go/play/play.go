package play

import (
	"fmt"
	"math/rand/v2"
	"time"

	"elem.com/roulette/game"
	"elem.com/roulette/halt"
	"elem.com/roulette/robot"
)

func RunTerminal(g *game.GameState) {
	ch := make(chan []int)

	go run(g, ch)

	for range ch {
	}
}

func Play() {
	window := robot.Window{}
	window.CaptureSize()
	window.SetReadyBarPosition(0)
	window.SetNumberArea()
	window.CaptureTerminal()

	rouletteMap, err := robot.UseRouletteMap("roulette.json", &window)
	if err != nil {
		fmt.Println("Error loading roulette map:", err)
		return
	}

	maxProtection := requestProtection()
	gState := game.NewGameState(maxProtection)

	numberArea := game.NewNumberArea(window.NumberArea)

	targetCh := make(chan []int)
	numCh := make(chan int)
	go runRobot(gState, targetCh, &window, numCh)

	doneCh := make(chan struct{})

	// Listen for the first number
	go numberArea.ReadNumber(numCh)

	for targets := range targetCh {
		halt.IsHalted.Store(false)

		go selectTargets(targets, &window, rouletteMap, doneCh)
		halt.ListenForHalt()

		<-doneCh

		go numberArea.ReadNumber(numCh)
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

func selectTargets(targets []int, window *robot.Window, rouletteMap *robot.RouletteMap, doneTargetCh chan struct{}) {
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

	doneTargetCh <- struct{}{}
}
