package main

import "fmt"

var RouletteNumbers = []int{
	10, 5, 24, 16, 33, 1, 20, 14, 31, 9, 22, 18, 29, 7, 28, 12, 35,
	3, 26, 0, 32, 15, 19, 4, 21, 2, 25, 17, 34, 6, 27, 13, 36,
	11, 30, 8, 23,
}

// RouletteNumberToIndex maps roulette numbers to their index in the RouletteNumbers slice
var RouletteNumberToIndex map[int]int

func init() {
	// Initialize the map
	RouletteNumberToIndex = make(map[int]int)
	for i, num := range RouletteNumbers {
		RouletteNumberToIndex[num] = i
	}
}

func main() {
	init()
	RunSimulations()
}

func CompareWithPrevious(drawn int, game []int) bool {
	for _, expected := range game {
		index := RouletteNumberToIndex[expected]

		target := RouletteNumbers[index]

		var prev, next int

		prev = RouletteNumbers[func() int {
			if index > 1 {
				return index - 1
			}
			return len(RouletteNumbers) - 1
		}()]

		next = RouletteNumbers[func() int {
			if index < len(RouletteNumbers)-1 {
				return index + 1
			}
			return 0
		}()]

		if prev == drawn || target == drawn || next == drawn {
			return true
		}
	}
	return false
}

func Play(numbers []int, protection int) []bool {
	curGame := []int{}
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

		if wonPrev || usedCount >= protection {
			expected, err := GetExpectedFor(previous)
			if err != nil {
				fmt.Printf("Warning: %v\n", err)
				continue
			}
			curGame = expected
		}

		win := CompareWithPrevious(drawn, curGame)

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

func GetMoney(numbers []int, bet float64, protection int, martin bool) {
	results := Play(numbers, protection)

	investment := 0.0

	pow := 0

	worstLoss := 0.0
	lastLostIndex := -1
	lastLostSeq := 0
	lostSeq := 0

	matches := 0

	for i, result := range results {
		if result {
			investment += bet * 36 * float64(1<<pow)
			matches++
		} else {
			investment -= bet * 12 * float64(1<<pow)
		}

		if investment < worstLoss {
			worstLoss = investment
		}

		if result {
			lastLostSeq = 0
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
		}
	}

	fmt.Printf("Out of %d numbers, %d numbers were found to be in the expected range\n", len(numbers), matches)
	fmt.Printf("You would have made R$ %.2f\n", investment)
	fmt.Printf("Worst case scenario: R$ %.2f\n", worstLoss)
	fmt.Printf("Longest lost sequence: %d\n", lostSeq)
	fmt.Printf("---\n\n")
}

func RunSimulations() {
	sliceSize := 25
	for i := 0; i < 20; i++ {
		start := i * sliceSize
		end := start + sliceSize
		if end > len(Last) {
			break
		}
		GetMoney(Last[start:end], 5, 2)
	}

	GetMoney(Last, 10, 2)
	GetMoney(Last, 10, 2)
}
