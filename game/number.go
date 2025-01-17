package game

import (
	"errors"
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"elem.com/roulette/halt"
	"elem.com/roulette/utils"

	"github.com/go-vgo/robotgo"
	"github.com/otiai10/gosseract/v2"
)

var (
	ErrNoNumber   = errors.New("no number")
	ErrWrongColor = errors.New("wrong color")
)

const (
	pointsNum     = 20
	pointsPerSide = 5
)

var (
	// Initialize Tesseract client
	client *gosseract.Client

	// Black numbers
	blackNumbers = []int{2, 4, 6, 8, 10, 11, 13, 15, 17, 20, 22, 24, 26, 28, 29, 31, 33, 35}
	isBlackArr   = [37]bool{}
)

type Color int

const (
	ColorRed Color = iota
	ColorGreen
	ColorBlack
)

type DrawnArea struct {
	bounds image.Rectangle
	points []image.Point
}

func init() {
	client = gosseract.NewClient()
	client.SetWhitelist("0123456789")

	for _, number := range blackNumbers {
		isBlackArr[number] = true
	}

	// Ensure directory exists
	if err := os.MkdirAll(utils.TmpDir, 0755); err != nil {
		log.Fatalf("error creating tmp directory: %v", err)
	}
}

func NewDrawnAreas(numBounds image.Rectangle, winBounds image.Rectangle) (numberArea *DrawnArea, winArea *DrawnArea) {
	numberArea = buildDrawArea(numBounds)
	winArea = buildDrawArea(winBounds)

	return
}

func buildDrawArea(bounds image.Rectangle) *DrawnArea {
	width := bounds.Dx()
	height := bounds.Dy()

	points := make([]image.Point, pointsNum)

	// Sides
	for side := 0; side < 4; side++ {
		for p := 0; p < pointsPerSide; p++ {
			var x, y int

			if side == 0 { // Top
				x = (width * p / pointsPerSide)
				y = 0
			} else if side == 1 { // Right
				x = width - 1
				y = (height * p / pointsPerSide)
			} else if side == 2 { // Bottom
				x = (width * p / pointsPerSide)
				y = height - 1
			} else if side == 3 { // Left
				x = 0
				y = (height * p / pointsPerSide)
			}

			points[side*pointsPerSide+p] = image.Point{X: x, Y: y}
		}
	}

	return &DrawnArea{
		bounds: bounds,
		points: points,
	}
}

func ReadNumber(ch chan int, numberArea *DrawnArea, winArea *DrawnArea) (found bool) {
	for {
		if halt.IsHalted() {
			utils.Console.Trace().Msgf("ReadNumber halted, waiting for continue...")
			time.Sleep(2 * time.Second)
			break
		}

		time.Sleep(500 * time.Millisecond)

		number, err := numberArea.captureNumber()
		if err != nil {
			handleFailNumber(err, "number area")

			number, err = winArea.captureNumber()
			if err != nil {
				handleFailNumber(err, "win area")
				continue
			}
		}

		utils.Console.Debug().Msgf("Read number: %d", number)

		ch <- number
		return true
	}

	return false
}

func handleFailNumber(err error, area string) {
	if errors.Is(err, ErrNoNumber) {
		utils.Console.Trace().Msgf("- [%s]", area)
		return
	}

	if errors.Is(err, ErrWrongColor) {
		utils.Console.Err(err).Msg("error wrong color")
		return
	}

	utils.Console.Err(err).Msg("error capturing number")
}

// CaptureNumber captures a screenshot of the specified region and performs OCR to extract a number
func (n *DrawnArea) captureNumber() (int, error) {
	img, err := robotgo.CaptureImg(
		n.bounds.Min.X,
		n.bounds.Min.Y,
		n.bounds.Dx(),
		n.bounds.Dy())
	if err != nil {
		return 0, fmt.Errorf("error capturing image: %w", err)
	}

	color, err := n.getColor(img)
	if err != nil {
		return 0, fmt.Errorf("error getting color: %w", err)
	}

	utils.Console.Trace().Msgf("Color: %d", color)

	procImg := processImage(img, color)

	return extractNumber(procImg, color)
}

func (n *DrawnArea) getColor(img image.Image) (Color, error) {
	// Initialize color to first pixel
	var color [3]uint32
	color[0], color[1], color[2], _ = img.At(0, 0).RGBA()

	// Compare with some tolerance as colors might not be exact
	const toleranceBlack uint32 = 0x1000 // 4096
	var isRed, isBlack, isGreen bool

	if color[0] > color[1] && color[0] > color[2] {
		isRed = true
	} else if color[1] > color[0] && color[1] > color[2] {
		isGreen = true
	} else if color[0] < toleranceBlack && color[1] < toleranceBlack && color[2] < toleranceBlack {
		isBlack = true
	}

	// If none found, it's not a number
	if !isRed && !isBlack && !isGreen {
		return 0, ErrNoNumber
	}

	// Check the borders for the same color
	for _, point := range n.points {
		r, g, b, _ := img.At(point.X, point.Y).RGBA()

		// The border color should be the same as the first pixel to be a number
		if color[0] != r || color[1] != g || color[2] != b {
			return 0, ErrNoNumber
		}
	}

	if isRed {
		return ColorRed, nil
	}

	if isGreen {
		return ColorGreen, nil
	}

	return ColorBlack, nil
}

func processImage(img image.Image, color Color) image.Image {
	if color == ColorBlack {
		// Inverted make the number more readable for OCR (black on white)
		img = utils.ProcessBlack(img)
	}

	if color == ColorRed {
		// Turn it black on white background
		img = utils.ProcessRed(img)
	}

	return img
}

func extractNumber(img image.Image, color Color) (int, error) {
	utils.Console.Trace().Msg("Extracting number")

	// Only 0 is green
	if color == ColorGreen {
		return 0, nil
	}

	// Tesseract doesn't work with []bits, I don't know why... Saving it to jpeg works.
	// TODO save png?
	robotgo.SaveJpeg(img, utils.TmpVerifyImg)
	if err := client.SetImage(utils.TmpVerifyImg); err != nil {
		return 0, fmt.Errorf("error setting image: %w", err)
	}

	text, err := client.Text()
	if err != nil {
		return 0, fmt.Errorf("error performing OCR: %w", err)
	}

	utils.Console.Trace().Msgf("OCR text: %s", text)

	number, err := strconv.Atoi(text)
	if err != nil {
		return handleFailValue(img, text, "error parsing number")
	}

	utils.Console.Trace().Msgf("Parsed number: %d", number)

	if err := validateNumber(number, color); err != nil {
		return handleFailValue(img, text, "error validating number")
	}

	if err := saveNumber(img, utils.ResultImgFolder, text); err != nil {
		utils.Console.Err(err).Msg("error saving number")
	}

	return number, nil
}

func handleFailValue(img image.Image, text string, errText string) (int, error) {
	if err := saveNumber(img, utils.FailedImgFolder, text); err != nil {
		utils.Console.Err(err).Msg("error saving number")
	}

	return 0, fmt.Errorf("%s: %w", errText, ErrNoNumber)
}

// Validates the number matches the expected color
func validateNumber(number int, color Color) error {
	if color == ColorBlack && !isBlackArr[number] {
		return fmt.Errorf("number should be BLACK but recognized is RED (%d): %w", number, ErrWrongColor)
	} else if color == ColorRed && isBlackArr[number] {
		return fmt.Errorf("number should be RED but recognized is BLACK (%d): %w", number, ErrWrongColor)
	}

	return nil
}

func saveNumber(img image.Image, folder string, number string) error {
	// Ensure directory exists
	dir := fmt.Sprintf("%s/%s", utils.OutputDir, folder)
	if err := os.MkdirAll(dir, 0755); err != nil {
		utils.Console.Fatal().Msgf("error creating output directory: %v", err)
	}

	randStr := fmt.Sprintf("%x", rand.Int31())
	filename := fmt.Sprintf("%d_%s_%s.png", time.Now().UnixMilli(), number, randStr)

	if err := robotgo.SaveJpeg(img, fmt.Sprintf("%s/%s", dir, filename)); err != nil {
		return fmt.Errorf("error saving number: %w", err)
	}

	return nil
}
