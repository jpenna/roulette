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
			fmt.Println("Por favor, insira um número válido")
			continue
		}

		g.computeWinsAndLosses()

		err = g.getBets()
		if err != nil {
			fmt.Println("Por favor, insira um número válido")
			continue
		}

		g.printTargets()

		ch <- g.targets
	}
}
