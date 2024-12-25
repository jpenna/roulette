package game

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"elem.com/roulette/halt"
	"elem.com/roulette/roulette"
	"elem.com/roulette/utils"
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

func NewGameState(maxProtection int) *GameState {
	return &GameState{
		maxProtection: maxProtection,
	}
}

func (g *GameState) GetTargets() []int {
	return g.targets
}

func (g *GameState) WaitDrawn(numCh chan int, next chan<- struct{}, endBet <-chan struct{}) {
	inputCh := make(chan string)
	errCh := make(chan error)

	go g.listenInput(inputCh, errCh)

	go func() {
		for {
			fmt.Println("Último número sorteado (#, [p]rint, [u]pdate, [s]top, [c]ontinue): ")
			<-endBet
		}
	}()

	for {
		// Wait for either input or OCR
		select {
		case input := <-inputCh:
			isNumber := g.handleInput(input)

			if isNumber {
				halt.Stop() // Stop processing to use the new number

				go func() {
					numCh <- g.lastDrawn
				}()
			}

		case err := <-errCh:
			utils.Console.Err(err).Msg("error reading input")

		case num := <-numCh:
			halt.Continue() // If a number is received, the game is resumed

			utils.Console.Debug().Msgf("Detected drawn (or input): %d", num)

			fmt.Println(num)
			g.lastDrawn = num

			next <- struct{}{}
		}
	}
}

func (g *GameState) listenInput(inputCh chan<- string, errCh chan error) {
	for {
		utils.Console.Debug().Msg("Adding input listener")

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			errCh <- fmt.Errorf("error reading input: %w", err)
			return
		}

		utils.Console.Trace().Msgf("Input: %s", input)

		inputCh <- strings.TrimSpace(input)

		utils.Console.Debug().Msg("Removing input listener")
	}
}

func (g *GameState) handleInput(input string) (isNumber bool) {
	switch input {
	case "u":
		utils.Console.Debug().Msg("Updating settings")

		err := g.UpdateSettings()
		if err != nil {
			utils.Console.Err(err).Msg("error updating settings")
		}

	case "p":
		utils.Console.Debug().Msg("Printing full game state")

		fmt.Println()
		g.PrintFullGameState()

	case "s":
		utils.Console.Trace().Msgf("Stopping game")

		halt.Stop()

	case "c":
		utils.Console.Trace().Msgf("Continuing game")

		halt.Continue()

	default:
		num, err := strconv.Atoi(input)
		if err != nil {
			utils.Console.Err(err).Msg("error getting number")
			return false
		}

		utils.Console.Debug().Msgf("Input drawn: %d", num)

		g.lastDrawn = num
		return true
	}

	return false
}

func (g *GameState) UpdateSettings() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n\n---\nMáximo de proteção (atual: %d): ", g.maxProtection)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("error getting number: %w", err)
	}

	g.maxProtection = num

	fmt.Printf("Proteção atualizada para %d\n---\n\n", g.maxProtection)

	return nil
}

func (g *GameState) ComputeWinsAndLosses() {
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
		fmt.Printf("\033[32m%s\033[0m\n\n---\n\n", message)
	} else {
		fmt.Printf("\033[31m%s\033[0m\n\n---\n\n", message)
	}
}

func (g *GameState) GetBets() error {
	// If it's the first game or we won last or protection is above max, get all bets
	if g.gameNumber == 0 || (g.wonLast || g.protectionCount >= g.maxProtection) {
		targets, bets, err := roulette.GetAllBetsFor(g.lastDrawn)
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

func (g *GameState) PrintTargets() {
	slices.SortFunc(g.targets, func(a, b int) int {
		return roulette.RouletteNumberToIndex[a] - roulette.RouletteNumberToIndex[b]
	})

	if g.usingProtection {
		fmt.Printf("Usando proteção (%d / %d)\n", g.protectionCount, g.maxProtection)
		fmt.Printf("\033[45mRepetir última aposta!\033[0m\n\n")
	} else {
		fmt.Printf("Alvos para %d:\n\n", g.lastDrawn)

		for _, t := range g.targets {
			fmt.Printf("  %d (posição: %d)\n", t, roulette.RouletteNumberToIndex[t]+1)
		}

		fmt.Println()
	}
}

func (g *GameState) PrintFullGameState() {
	fmt.Printf("\033[47m\033[30mJogo %d: %d vitórias, %d derrotas (%.1f%%)\033[0m\n", g.gameNumber, g.wins, g.losses, float32(g.wins*100)/float32(g.gameNumber))
	fmt.Printf("\033[47m\033[30mProteção: %d / %d\033[0m\n", g.protectionCount, g.maxProtection)
	fmt.Printf("\033[47m\033[30mProteção ativa: %t\033[0m\n", g.usingProtection)
	fmt.Printf("\033[47m\033[30mÚltimo sorteado: %d\033[0m\n", g.lastDrawn)
	fmt.Printf("\033[47m\033[30mAlvos: %v\033[0m\n", g.targets)
	fmt.Printf("\033[47m\033[30mApostas: %v\033[0m\n", g.bets)
	fmt.Println()
}
