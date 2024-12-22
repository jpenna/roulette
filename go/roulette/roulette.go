package roulette

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
