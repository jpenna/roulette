package simulation

import (
	"fmt"

	"elem.com/roulette/roulette"
	"elem.com/roulette/utils"
)

func Run(numbers []int, bet float64, protection int) float64 {
	return execute(numbers, bet, protection, false)
}

func RunMartingale(numbers []int, bet float64, protection int) float64 {
	return execute(numbers, bet, protection, true)
}

func execute(numbers []int, bet float64, protection int, martin bool) float64 {
	results := play(numbers, protection)

	investment := 0.0

	pow := 0

	worstLoss := 0.0
	worstLossIndex := -1
	lastLostIndex := -1
	lastLostSeq := 0
	lostSeq := 0
	lostSeqIndex := -1

	maxWin := 0.0
	maxWinIndex := -1

	matches := 0

	for i, win := range results {
		gameCost := bet * 12 * float64(uint(1)<<pow)

		if win {
			investment += bet*36*float64(uint(1)<<pow) - gameCost
			matches++
		} else {
			investment -= gameCost
		}

		if investment < worstLoss {
			worstLoss = investment
			worstLossIndex = i
		}

		if win {
			lastLostSeq = 0
			pow = 0
		} else {
			if i == lastLostIndex+1 {
				lastLostSeq++
			}
			lastLostIndex = i

			if martin {
				pow++
			}
		}

		if lastLostSeq > lostSeq {
			lostSeq = lastLostSeq
			lostSeqIndex = i
		}

		if investment > maxWin {
			maxWin = investment
			maxWinIndex = i
		}
	}

	fmt.Printf("De %d números, %d números foram sorteados (%d%%)\n", len(numbers), matches, (matches*100)/len(numbers))
	if investment > 0 {
		fmt.Printf("\033[32mVocê teria ganho R$ %.2f\033[0m\n", investment)
	} else {
		fmt.Printf("\033[31mVocê teria perdido R$ %.2f\033[0m\n", investment)
	}

	fmt.Printf("Melhor cenário possível: R$ %.2f (%d)\n", maxWin, maxWinIndex+1)
	fmt.Printf("Pior cenário possível: R$ %.2f (%d)\n", worstLoss, worstLossIndex+1)

	fmt.Printf("Maior sequência de perdas: %d (%d)\n", lostSeq, lostSeqIndex+1)
	fmt.Printf("---\n\n")

	return investment
}

func play(numbers []int, protection int) []bool {
	curBets := []int{}
	wonPrev := true
	usedCount := 0

	// Create a reversed copy of numbers
	reversed := make([]int, len(numbers))
	copy(reversed, numbers)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	// Remove last number and process the rest
	results := make([]bool, len(reversed)-1)
	for i := 0; i < len(reversed)-1; i++ {
		previous := reversed[i]
		drawn := reversed[i+1]

		if wonPrev || usedCount > protection {
			_, bets, err := roulette.GetAllBetsFor(previous)
			if err != nil {
				utils.Console.Warn().Err(err).Msg("warning: error getting all bets")
				continue
			}
			curBets = bets
		}

		win := checkWin(drawn, curBets)

		if win {
			usedCount = 0
			wonPrev = true
			results[i] = true
		} else {
			usedCount++
			wonPrev = false
			results[i] = false
		}
	}

	return results
}

func checkWin(drawn int, bets []int) bool {
	for _, bet := range bets {
		if bet == drawn {
			return true
		}
	}
	return false
}
