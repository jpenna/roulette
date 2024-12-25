package robot

// 1. Get window region
// 2. Build an Area struct containing the 4 corners of each number in the roulette, one by one, asking for the number and each corner and waiting for the user to press enter with the pointer over each corner
// 3. Based on the map and window region, compute new values for the corners considering the topLeft of the window 0, 0
// 4. store the corner values in a map for the number
// 5. store the map in a file that can be read by the program (also store the window region)

// 6. Create a function that computes where the number is located relative to the window region, considering the window position and size can change

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"

	"elem.com/roulette/roulette"
	"elem.com/roulette/utils"
	"github.com/go-vgo/robotgo"
)

// Area represents a rectangular region defined by 4 corners
type Area struct {
	TopLeft     [2]int `json:"top_left"`
	TopRight    [2]int `json:"top_right"`
	BottomLeft  [2]int `json:"bottom_left"`
	BottomRight [2]int `json:"bottom_right"`
}

// RouletteMap stores the mapping of numbers to their areas on the screen
type RouletteMap struct {
	WindowRegion *Window      `json:"windowRegion"`
	NumberAreas  map[int]Area `json:"numberAreas"`
}

func UseRouletteMap(filename string, window *Window) (*RouletteMap, error) {
	absolutePath := buildAbsoluteFilePath(filename)

	rouletteMap, err := loadRouletteMap(absolutePath)
	if err != nil {
		return nil, fmt.Errorf("error loading roulette map: %w", err)
	}

	rouletteMap.adjustCoordinates(window)

	return rouletteMap, nil
}

func NewRouletteMap(filename string) {
	absolutePath := buildAbsoluteFilePath(filename)

	rouletteMap, err := buildRouletteMap(absolutePath)
	if err != nil {
		utils.Console.Err(err).Msg("error building roulette map")
		return
	}

	rouletteMap.saveToFile(absolutePath)

	utils.Console.Info().Msgf("Roulette map saved to %s", filename)
}

// BuildRouletteMap creates a new roulette map by capturing the window region
// and the coordinates for each number on the roulette wheel
func buildRouletteMap(filename string) (*RouletteMap, error) {
	// First capture the window region
	window := &Window{}
	window.CaptureSize()

	// Initialize the map
	rouletteMap := &RouletteMap{
		WindowRegion: window,
		NumberAreas:  make(map[int]Area),
	}

	reader := bufio.NewReader(os.Stdin)

	// For each number 0-36, capture the coordinates
	for i := 0; i < len(roulette.RouletteNumbers); i++ {
		num := roulette.RouletteNumbers[i]
		fmt.Printf("\nCapturing coordinates for number %d\n", num)
		area := captureArea(window, reader)
		rouletteMap.NumberAreas[num] = area

		rouletteMap.saveToFile(filename)
	}

	return rouletteMap, nil
}

// captureArea captures the 4 corners of an area from user input
func captureArea(window *Window, reader *bufio.Reader) Area {
	area := Area{}

	fmt.Println("Position mouse at top-left corner and press Enter")
	reader.ReadString('\n')
	x, y := robotgo.Location()
	area.TopLeft[0] = x - window.TopLeft[0]
	area.TopLeft[1] = y - window.TopLeft[1]

	fmt.Println("Position mouse at top-right corner and press Enter")
	reader.ReadString('\n')
	x, y = robotgo.Location()
	area.TopRight[0] = x - window.TopLeft[0]
	area.TopRight[1] = y - window.TopLeft[1]

	fmt.Println("Position mouse at bottom-right corner and press Enter")
	reader.ReadString('\n')
	x, y = robotgo.Location()
	area.BottomRight[0] = x - window.TopLeft[0]
	area.BottomRight[1] = y - window.TopLeft[1]

	fmt.Println("Position mouse at bottom-left corner and press Enter")
	reader.ReadString('\n')
	x, y = robotgo.Location()
	area.BottomLeft[0] = x - window.TopLeft[0]
	area.BottomLeft[1] = y - window.TopLeft[1]

	return area
}

// SaveToFile saves the roulette map to a JSON file
func (rm *RouletteMap) saveToFile(filename string) error {
	data, err := json.MarshalIndent(rm, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling roulette map: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing roulette map file: %w", err)
	}

	return nil
}

// LoadFromFile loads a roulette map from a JSON file
func loadRouletteMap(filename string) (*RouletteMap, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading roulette map file: %w", err)
	}

	var rm RouletteMap
	err = json.Unmarshal(data, &rm)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling roulette map: %w", err)
	}

	return &rm, nil
}

// AdjustCoordinates returns the relative coordinates of a number based on the current window position
func (rm *RouletteMap) adjustCoordinates(window *Window) error {
	for number := range rm.NumberAreas {
		adjustedArea, err := rm.adjustCoordinatesFor(window, number)
		if err != nil {
			return err
		}
		rm.NumberAreas[number] = *adjustedArea
	}

	return nil
}

func (rm *RouletteMap) adjustCoordinatesFor(window *Window, number int) (*Area, error) {
	area, exists := rm.NumberAreas[number]
	if !exists {
		return nil, fmt.Errorf("number %d not found in roulette map", number)
	}

	// Calculate the offset between the original window position and current window position
	offsetX := float64(window.TopLeft[0])
	offsetY := float64(window.TopLeft[1])

	scaleX := float64(window.BottomRight[0]-window.TopLeft[0]) / float64(rm.WindowRegion.BottomRight[0]-rm.WindowRegion.TopLeft[0])
	scaleY := float64(window.BottomRight[1]-window.TopLeft[1]) / float64(rm.WindowRegion.BottomRight[1]-rm.WindowRegion.TopLeft[1])

	// Apply the offset to all corners
	adjustedArea := Area{
		TopLeft:     [2]int{int(scaleX*float64(area.TopLeft[0]) + offsetX), int(scaleY*float64(area.TopLeft[1]) + offsetY)},
		TopRight:    [2]int{int(scaleX*float64(area.TopRight[0]) + offsetX), int(scaleY*float64(area.TopRight[1]) + offsetY)},
		BottomLeft:  [2]int{int(scaleX*float64(area.BottomLeft[0]) + offsetX), int(scaleY*float64(area.BottomLeft[1]) + offsetY)},
		BottomRight: [2]int{int(scaleX*float64(area.BottomRight[0]) + offsetX), int(scaleY*float64(area.BottomRight[1]) + offsetY)},
	}

	return &adjustedArea, nil
}

func buildAbsoluteFilePath(filename string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		return filename
	}
	return path.Join(currentDir, "robot", "maps", filename)
}

func (rm *RouletteMap) ClickNumber(number int) {
	area, exists := rm.NumberAreas[number]
	if !exists {
		return
	}

	// Generate two random values between 0 and 1
	s := rand.Float64()
	t := rand.Float64()

	// Ensure s + t <= 1 for proper distribution
	if s+t > 1 {
		s = 1 - s
		t = 1 - t
	}

	// Bilinear interpolation using all four corners
	x := int(float64(area.TopLeft[0])*(1-s)*(1-t) +
		float64(area.TopRight[0])*s*(1-t) +
		float64(area.BottomLeft[0])*(1-s)*t +
		float64(area.BottomRight[0])*s*t)

	y := int(float64(area.TopLeft[1])*(1-s)*(1-t) +
		float64(area.TopRight[1])*s*(1-t) +
		float64(area.BottomLeft[1])*(1-s)*t +
		float64(area.BottomRight[1])*s*t)

	utils.Console.Trace().Msgf("Clicking number %d at coordinates: %d, %d", number, x, y)
	Click(x, y)
}
