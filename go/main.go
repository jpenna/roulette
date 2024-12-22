package main

import (
	"fmt"
	"os"

	"elem.com/roulette/simulation"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("No arguments provided. Use --simulate or --play")
	}

	switch args[0] {
	case "--simulate":
		RunSimulations()
		return
	case "--play":
		fmt.Println("playing")
		return
	default:
		panic("Invalid argument. Use --simulate or --play")
	}
}

func RunSimulations() {
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

// func printBetsFor(x int) {
// 	bets, err := game.GetAllBets(x)
// 	if err != nil {
// 		log.Println("Error getting bets for ", x, ":", err)
// 	}

// 	sort.Ints(bets)

// 	log.Println("Bets for ", x, ":", bets)
// }
