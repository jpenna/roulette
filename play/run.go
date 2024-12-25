package play

import (
	"fmt"

	"elem.com/roulette/game"
)

func run(g *game.GameState, targetCh chan<- []int, numCh chan int) {
	for {
		err := g.RequestNumber(numCh)
		if err != nil {
			fmt.Print("\033[41m")
			fmt.Print("Por favor, insira um número válido")
			fmt.Print("\033[0m\n")
			continue
		}

		g.ComputeWinsAndLosses()

		err = g.GetBets()
		if err != nil {
			fmt.Print("\033[41m")
			fmt.Print("Por favor, insira um número válido")
			fmt.Print("\033[0m\n")
			continue
		}

		g.PrintTargets()

		targetCh <- g.GetTargets()
	}
}
