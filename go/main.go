package main

import (
	"log"
	"os"

	"elem.com/roulette/play"
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
		play.RunTerminal()
		return
	case "--bets":
		log.SetFlags(0)
		printBets()
	default:
		panic("Invalid argument. Use --simulate or --play")
	}
}

func runSimulations() {
	sliceSize := 25
	for i := 0; i < 20; i++ {
		start := i * sliceSize
		end := start + sliceSize
		if end > len(simulation.Last) {
			break
		}
		simulation.Run(simulation.Last[start:end], 5, 2)
	}

	simulation.Run(simulation.Last, 10, 2)
}
