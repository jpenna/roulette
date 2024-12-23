package play

import (
	"fmt"
	"time"

	"elem.com/roulette/robot"
)

func (g *GameState) RunTerminal() {
	ch := make(chan []int)

	go g.run(ch)

	for range ch {
	}
}

func (g *GameState) RunRobot(ch chan<- []int, window *robot.Window) {
	g.runAuto(ch, window)
}

func (g *GameState) runAuto(ch chan<- []int, window *robot.Window) {
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

		waitReady(window, ch)

		ch <- g.targets
	}
}

func waitReady(window *robot.Window, ch chan<- []int) {
	waited := 0

	for {
		ready, err := window.IsReadyToBet()
		if err != nil {
			fmt.Println("Erro ao verificar se está pronto para apostar: ", err)

			fmt.Print("Continuar? (y/n): ")
			var input string
			fmt.Scanln(&input)
			if input != "y" {
				return
			}

			close(ch)
			return
		}

		if ready {
			break
		}

		time.Sleep(500 * time.Millisecond)
		waited += 500

		if waited%10_000 == 0 {
			fmt.Printf("Aguardando liberação para apostar (%ds)...\n", waited/1000)
		}
	}
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
