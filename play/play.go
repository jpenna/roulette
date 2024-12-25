package play

import (
	"fmt"
	"math/rand/v2"
	"time"

	"elem.com/roulette/game"
	"elem.com/roulette/halt"
	"elem.com/roulette/robot"
	"elem.com/roulette/utils"
)

func Play() {
	window := robot.Window{}
	window.CaptureSize()
	window.SetReadyBarPosition(0)
	window.SetNumberAreas()
	window.CaptureTerminal()

	rouletteMap, err := robot.UseRouletteMap("roulette.json", &window)
	if err != nil {
		utils.Console.Err(err).Msg("error loading roulette map")
		return
	}

	maxProtection := requestProtection()
	gState := game.NewGameState(maxProtection)

	numberArea, winArea := game.NewDrawnAreas(window.NumberArea, window.WinArea)

	targetCh := make(chan []int)
	numCh := make(chan int)
	betCh := make(chan struct{})
	go runRobot(gState, targetCh, &window, numCh, betCh)

	// Listen for the first number
	go func() {
		for {
			found := game.ReadNumber(numCh, numberArea, winArea)
			if found {
				utils.Console.Debug().Msgf("Number found")
				// The number is read, wait so it won't read again while the game is running
				time.Sleep(20 * time.Second)
			}
		}
	}()

	for targets := range targetCh {
		selectTargets(targets, &window, rouletteMap)
		betCh <- struct{}{}
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
		if halt.IsHalted() {
			break
		}

		rouletteMap.ClickNumber(target)

		delay := time.Duration(300+rand.Float64()*1000) * time.Millisecond
		time.Sleep(delay)
	}

	if !halt.IsHalted() {
		window.ClickTerminal()
	}
}
