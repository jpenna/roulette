package play

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"elem.com/roulette/roulette"
)

func RunTerminal() {
	gameNumber := 1

	for {
		fmt.Printf("Game %d\n", gameNumber)
		gameNumber++

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a number: ")
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number")
		}

		targets, err := roulette.GetTargetBetsFor(num)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		slices.SortFunc(targets, func(a, b int) int {
			return roulette.RouletteNumberToIndex[a] - roulette.RouletteNumberToIndex[b]
		})

		fmt.Printf("Expected numbers for %d:\n\n", num)

		for _, t := range targets {
			fmt.Printf("%d (position: %d)\n", t, roulette.RouletteNumberToIndex[t]+1)
		}

		fmt.Println()
	}
}
