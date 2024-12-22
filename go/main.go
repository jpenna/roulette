package main

import (
	"log"
	"os"

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
		play.NewGameState().RunTerminal()
		return
	case "--bets":
		log.SetFlags(0)
		printBets()
	case "--robot":
		robot.Screen()
	default:
		panic("Invalid argument. Use --simulate or --play")
	}
}

func runSimulations() {
	// sliceSize := 25
	// for i := 0; i < 20; i++ {
	// 	start := i * sliceSize
	// 	end := start + sliceSize
	// 	if end > len(simulation.Last) {
	// 		break
	// 	}
	// 	simulation.Run(simulation.Last[start:end], 5, 1)
	// }

	simulation.Run(simulation.Last4[0:100], 1, 2)
	simulation.Run(simulation.Last4, 1, 2)
}
