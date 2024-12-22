package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"elem.com/roulette/game"
)

func printBets() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a number: ")
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil {
			log.Println("Please enter a valid number")
			continue
		}

		printBetsFor(num)
	}
}

func printBetsFor(x int) {
	targets, bets, err := game.GetAllBets(x)
	if err != nil {
		log.Println("Error getting bets for ", x, ":", err)
	}

	sort.Ints(bets)

	log.Printf("\nTargets for %d: %v", x, targets)
	log.Printf("Bets for %d: %v\n\n", x, bets)
}
