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
	"image"
	"image/color"
	"math/rand"
	"os"
	"path"

	"elem.com/roulette/roulette"
	"elem.com/roulette/utils"
	"github.com/go-vgo/robotgo"
)

// RouletteMap stores the mapping of numbers to their areas on the screen
type RouletteMap struct {
	WindowRegion *Window        `json:"windowRegion"`
	Middle       map[int][2]int `json:"middle"`

	MiddleAreas  map[int]image.Rectangle `json:"middleAreas"`
	PixelsToSide int                     `json:"pixelsToSide"`
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
		Middle:       make(map[int][2]int),
		MiddleAreas:  make(map[int]image.Rectangle),
	}

	reader := bufio.NewReader(os.Stdin)

	rouletteMap.PixelsToSide = capturePixelsToSide(reader)

	// For each number 0-36, capture the coordinates
	for i := 0; i < len(roulette.RouletteNumbers); i++ {
		num := roulette.RouletteNumbers[i]
		fmt.Printf("\nCapturing coordinates for number %d\n", num)
		middle := captureMiddle(window, reader)
		rouletteMap.Middle[num] = middle

		rouletteMap.saveToFile(filename)
	}

	return rouletteMap, nil
}

func capturePixelsToSide(reader *bufio.Reader) int {
	fmt.Println("Put mouse to the LEFT of 6 and press Enter")
	reader.ReadString('\n')
	xLeft, _ := robotgo.Location()

	fmt.Println("Put mouse to the RIGHT of 6 and press Enter")
	reader.ReadString('\n')
	xRight, _ := robotgo.Location()

	return xRight - xLeft
}

// captureMiddle captures the middle of the number
func captureMiddle(window *Window, reader *bufio.Reader) [2]int {
	fmt.Println("Position mouse in the middle and press Enter")
	reader.ReadString('\n')
	x, y := robotgo.Location()
	return [2]int{x - window.TopLeft[0], y - window.TopLeft[1]}
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
	// Calculate the offset between the original window position and current window position
	offsetX := window.TopLeft[0]
	offsetY := window.TopLeft[1]

	scaleX := float64(window.BottomRight[0]-window.TopLeft[0]) / float64(rm.WindowRegion.BottomRight[0]-rm.WindowRegion.TopLeft[0])
	scaleY := float64(window.BottomRight[1]-window.TopLeft[1]) / float64(rm.WindowRegion.BottomRight[1]-rm.WindowRegion.TopLeft[1])

	pixelsToSide := float64(rm.PixelsToSide) * scaleX / 2

	for number := range rm.Middle {
		middle, exists := rm.Middle[number]
		if !exists {
			return fmt.Errorf("number %d not found in roulette map", number)
		}

		// Apply the offset to all corners
		rm.MiddleAreas[number] = image.Rectangle{
			Min: image.Point{int(scaleX*float64(middle[0]+offsetX) - pixelsToSide), int(scaleY*float64(middle[1]+offsetY) - pixelsToSide)},
			Max: image.Point{int(scaleX*float64(middle[0]+offsetX) + pixelsToSide), int(scaleY*float64(middle[1]+offsetY) + pixelsToSide)},
		}
	}

	return nil
}

func buildAbsoluteFilePath(filename string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		return filename
	}
	return path.Join(currentDir, "robot", "maps", filename)
}

func (rm *RouletteMap) ClickNumber(number int) {
	area, exists := rm.MiddleAreas[number]
	if !exists {
		return
	}

	x := area.Min.X + rand.Intn(area.Max.X-area.Min.X+1)
	y := area.Min.Y + rand.Intn(area.Max.Y-area.Min.Y+1)

	utils.Console.Trace().Msgf("Clicking number %d at coordinates: %d, %d", number, x, y)
	Click(x, y)
}

func (rm *RouletteMap) PrintMap(window *Window) *image.Rectangle {
	// Take a screenshot of the entire window area
	width := window.BottomRight[0] - window.TopLeft[0]
	height := window.BottomRight[1] - window.TopLeft[1]

	img, err := robotgo.CaptureImg(window.TopLeft[0], window.TopLeft[1], width, height)
	if err != nil {
		utils.Console.Err(err).Msg("Failed to capture screenshot")
		return nil
	}

	// Create a new RGBA image to draw on
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	// Copy the original image
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	// Draw cyan borders around each number area
	cyan := color.RGBA{0, 255, 255, 255} // Cyan color
	for _, area := range rm.MiddleAreas {
		// Adjust coordinates relative to the window
		relativeArea := image.Rect(
			area.Min.X-window.TopLeft[0],
			area.Min.Y-window.TopLeft[1],
			area.Max.X-window.TopLeft[0],
			area.Max.Y-window.TopLeft[1],
		)

		// Draw filled rectangle
		for x := relativeArea.Min.X; x <= relativeArea.Max.X; x++ {
			if x >= 0 && x < bounds.Max.X {
				for y := relativeArea.Min.Y; y <= relativeArea.Max.Y; y++ {
					if y >= 0 && y < bounds.Max.Y {
						rgba.Set(x, y, cyan)
					}
				}
			}
		}
	}

	// Save the image
	if err := robotgo.SaveJpeg(rgba, fmt.Sprintf("%s/roulette_map.jpg", utils.DataDir)); err != nil {
		utils.Console.Error().Err(err).Msg("Failed to save image")
	}

	return nil
}
