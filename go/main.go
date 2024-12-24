package main

import (
	"fmt"
	"log"
	"os"

	"elem.com/roulette/game"
	"elem.com/roulette/play"
	"elem.com/roulette/robot"
	"elem.com/roulette/simulation"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("No arguments provided. Use --simulate or --play")
	}

	switch args[0] {
	case "--simulate":
		log.SetFlags(0)
		runSimulations()
		return
	case "--play-terminal":
		log.SetFlags(0)
		gState := game.NewGameState(2)
		play.RunTerminal(gState)
		return
	case "--bets":
		log.SetFlags(0)
		printBets()
	case "--robot":
		play.Play()
	case "--build-map":
		robot.NewRouletteMap("roulette.json")
	case "--mouse-position":
		robot.MousePosition()
	case "--play-robot":
		log.SetFlags(0)
		log.Println("Iniciando bot...")
		log.Println("Confirme que o número vizinho é 1.")
		play.Play()
	case "--find-number":
		// logToFile("number.log")
		robot.MatchNumbers()
	default:
		panic("Invalid argument. Use --simulate or --play")
	}
}

func runSimulations() {
	list := simulation.Combined2
	chipValue := 2.5

	sum0 := 0.0
	sum1 := 0.0
	sum2 := 0.0

	sliceSize := 50
	for i := 0; i < len(list)/sliceSize; i++ {

		start := i * sliceSize
		end := start + sliceSize
		if end > len(list) {
			break
		}

		res0 := simulation.Run(list[start:end], chipValue, 0)
		res1 := simulation.Run(list[start:end], chipValue, 1)
		res2 := simulation.Run(list[start:end], chipValue, 2)

		fmt.Printf("0: %f\n", res0)
		fmt.Printf("1: %f\n", res1)
		fmt.Printf("2: %f\n", res2)
		fmt.Println()

		sum0 += res0
		sum1 += res1
		sum2 += res2

	}

	fmt.Printf("sum 0: %f\n", sum0)
	fmt.Printf("sum 1: %f\n", sum1)
	fmt.Printf("sum 2: %f\n", sum2)

	simulation.Run(list, chipValue, 0)
	simulation.Run(list, chipValue, 1)
	simulation.Run(list, chipValue, 2)
}
