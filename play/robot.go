package play

import (
	"fmt"
	"time"

	"elem.com/roulette/game"
	"elem.com/roulette/robot"
	"elem.com/roulette/utils"
)

func runRobot(g *game.GameState, targetCh chan<- []int, window *robot.Window, numCh chan int) {
	for {
		err := g.RequestNumber(numCh)
		if err != nil {
			fmt.Print("\033[41m")
			fmt.Print("Por favor, insira um número válido")
			fmt.Print("\033[0m\n")
			continue
		}

		g.ComputeWinsAndLosses()

		if err := g.GetBets(); err != nil {
			fmt.Print("\033[41m")
			fmt.Print("Por favor, insira um número válido")
			fmt.Print("\033[0m\n")
			continue
		}

		g.PrintTargets()

		waitReady(window, targetCh)

		targetCh <- g.GetTargets()
	}
}

func waitReady(window *robot.Window, ch chan<- []int) {
	waited := 0

	for {
		ready, err := game.IsReadyToBet(window.ReadyBarPosition[0], window.ReadyBarPosition[1])
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

		utils.Console.Trace().Msgf("ready: %t", ready)

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
