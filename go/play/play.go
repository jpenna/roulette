package play

import (
	"fmt"
)

func (g *GameState) RunTerminal() {
	ch := make(chan []int)

	go g.run(ch)

	for range ch {
	}
}

func (g *GameState) RunRobot(ch chan<- []int) {
	g.run(ch)
}

func (g *GameState) run(ch chan<- []int) {
	for {
		err := g.requestNumber()
		if err != nil {
			fmt.Print("\033[41m")
			fmt.Print("Por favor, insira um número válido")
			fmt.Print("\033[0m\n")
			continue
		}

		g.computeWinsAndLosses()

		err = g.getBets()
		if err != nil {
			fmt.Print("\033[41m")
			fmt.Print("Por favor, insira um número válido")
			fmt.Print("\033[0m\n")
			continue
		}

		g.printTargets()

		ch <- g.targets
	}
}
