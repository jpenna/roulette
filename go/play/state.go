package play

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"elem.com/roulette/game"
	"elem.com/roulette/roulette"
)

type GameState struct {
	gameNumber int

	wins   int
	losses int

	maxProtection int

	wonLast         bool
	protectionCount int
	usingProtection bool
	lastDrawn       int

	targets []int
	bets    []int
}

func NewGameState() *GameState {
	return &GameState{
		maxProtection: 0,
	}
}

func (g *GameState) requestNumber() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Último número sorteado: ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	if input == "u" {
		err := g.updateSettings()
		if err != nil {
			return fmt.Errorf("error updating settings: %w", err)
		}
		return g.requestNumber()
	}

	if input == "p" {
		log.Println()
		g.PrintFullGameState()
		return g.requestNumber()
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("error getting number: %w", err)
	}

	g.lastDrawn = num

	return nil
}

func (g *GameState) updateSettings() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n\n---\nMáximo de proteção (atual: %d): ", g.maxProtection)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("error getting number: %w", err)
	}

	g.maxProtection = num

	log.Printf("Proteção atualizada para %d\n---\n\n", g.maxProtection)

	return nil
}

func (g *GameState) computeWinsAndLosses() {
	if len(g.targets) == 0 {
		return
	}

	g.wonLast = false
	for _, bet := range g.bets {
		if bet == g.lastDrawn {
			g.wonLast = true
			break
		}
	}

	if g.wonLast {
		g.wins++
	} else {
		g.losses++
	}

	message := fmt.Sprintf("\nJogo %d: %d vitórias, %d derrotas (%.1f%%)", g.gameNumber, g.wins, g.losses, float32(g.wins*100)/float32(g.gameNumber))
	if g.wonLast {
		log.Printf("\033[32m%s\033[0m\n\n---\n\n", message)
	} else {
		log.Printf("\033[31m%s\033[0m\n\n---\n\n", message)
	}
}

func (g *GameState) getBets() error {
	// If it's the first game or we won last or protection is above max, get all bets
	if g.gameNumber == 0 || (g.wonLast || g.protectionCount >= g.maxProtection) {
		targets, bets, err := game.GetAllBets(g.lastDrawn)
		if err != nil {
			return fmt.Errorf("error getting bets: %w", err)
		}

		g.targets = targets
		g.bets = bets
		g.usingProtection = false
		g.protectionCount = 0
	} else {
		g.protectionCount++
		g.usingProtection = true
	}

	g.gameNumber++

	return nil
}

func (g *GameState) printTargets() {
	slices.SortFunc(g.targets, func(a, b int) int {
		return roulette.RouletteNumberToIndex[a] - roulette.RouletteNumberToIndex[b]
	})

	if g.usingProtection {
		log.Printf("Usando proteção (%d / %d)\n", g.protectionCount, g.maxProtection)
		log.Printf("\033[45mRepetir última aposta!\033[0m\n\n")
	} else {
		log.Printf("Alvos para %d:\n\n", g.lastDrawn)

		for _, t := range g.targets {
			log.Printf("  %d (posição: %d)\n", t, roulette.RouletteNumberToIndex[t]+1)
		}

		log.Println()
	}

}

func (g *GameState) PrintFullGameState() {
	log.Printf("\033[47m\033[30mJogo %d: %d vitórias, %d derrotas (%.1f%%)\033[0m\n", g.gameNumber, g.wins, g.losses, float32(g.wins*100)/float32(g.gameNumber))
	log.Printf("\033[47m\033[30mProteção: %d / %d\033[0m\n", g.protectionCount, g.maxProtection)
	log.Printf("\033[47m\033[30mProteção ativa: %t\033[0m\n", g.usingProtection)
	log.Printf("\033[47m\033[30mÚltimo sorteado: %d\033[0m\n", g.lastDrawn)
	log.Printf("\033[47m\033[30mAlvos: %v\033[0m\n", g.targets)
	log.Printf("\033[47m\033[30mApostas: %v\033[0m\n", g.bets)
	log.Println()
}
