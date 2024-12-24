package game

import (
	"errors"
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"
	"strconv"

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

	tmpDir       = "data/tmp"
	tmpVerifyImg = "data/tmp/capture.jpeg"
	outputDir    = "data/numbers"
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

type NumberArea struct {
	bounds image.Rectangle
	points []image.Point
}

func init() {
	client = gosseract.NewClient()
	client.SetWhitelist("0123456789")
	client.SetVariable("tessedit_write_images", "true")

	for _, number := range blackNumbers {
		isBlackArr[number] = true
	}

	// Ensure directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("error creating output directory: %v", err)
	}
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		log.Fatalf("error creating tmp directory: %v", err)
	}

}

func NewNumberArea(bounds image.Rectangle) *NumberArea {
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

	return &NumberArea{
		bounds: bounds,
		points: points,
	}
}

// CaptureNumber captures a screenshot of the specified region and performs OCR to extract a number
func (n *NumberArea) CaptureNumber() (int, error) {
	img, err := robotgo.CaptureImg(
		n.bounds.Min.X,
		n.bounds.Min.Y,
		n.bounds.Dx(),
		n.bounds.Dy())
	if err != nil {
		return 0, fmt.Errorf("error capturing image: %w", err)
	}

	if err := saveNumber(img, "capture", false, false); err != nil {
		fmt.Printf("error saving number: %v", err)
	}

	color, err := n.getColor(img)
	if err != nil {
		return 0, fmt.Errorf("error getting color: %w", err)
	}

	return processImage(img, color, false)
}

func (n *NumberArea) getColor(img image.Image) (Color, error) {
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

func processImage(img image.Image, color Color, invert bool) (int, error) {
	if invert {
		log.Println("Trying with inverted image")
		// Try inverted to make the number more readable for OCR
		img = invertImage(img)
	}

	// Only 0 is green
	if color == ColorGreen {
		return 0, nil
	}

	// Tesseract doesn't work with []bits, I don't know why... Saving it to jpeg works.
	robotgo.SaveJpeg(img, tmpVerifyImg)
	if err := client.SetImage(tmpVerifyImg); err != nil {
		return 0, fmt.Errorf("error setting image: %w", err)
	}

	text, err := client.Text()
	if err != nil {
		return 0, fmt.Errorf("error performing OCR: %w", err)
	}

	number, err := strconv.Atoi(text)
	if err != nil {
		return handleFailNumber(img, color, text, invert, "error parsing number")
	}

	if err := validateNumber(number, color); err != nil {
		return handleFailNumber(img, color, text, invert, "error validating number")
	}

	if err := saveNumber(img, text, invert, true); err != nil {
		fmt.Printf("error saving number: %v", err)
	}

	return number, nil
}

func handleFailNumber(img image.Image, color Color, text string, invert bool, errText string) (int, error) {
	if err := saveNumber(img, text, invert, false); err != nil {
		log.Printf("error saving number: %v", err)
	}
	if !invert {
		return processImage(img, color, true)
	}

	return 0, fmt.Errorf("%s: %w", errText, ErrNoNumber)
}

func invertImage(img image.Image) image.Image {
	bounds := img.Bounds()
	inverted := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			// Convert from 16-bit to 8-bit color components and invert
			i := (y-bounds.Min.Y)*inverted.Stride + (x-bounds.Min.X)*4
			inverted.Pix[i+0] = uint8(255 - (r >> 8)) // R
			inverted.Pix[i+1] = uint8(255 - (g >> 8)) // G
			inverted.Pix[i+2] = uint8(255 - (b >> 8)) // B
			inverted.Pix[i+3] = uint8(a >> 8)         // A
		}
	}

	return inverted
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

func saveNumber(img image.Image, number string, invert bool, valid bool) error {
	suffix := ""
	if invert {
		suffix = "inverted"
	}
	if !valid {
		suffix += "_invalid"
	}

	randStr := fmt.Sprintf("%x", rand.Int31())

	// Saving it to png makes it fully transparent, so using jpeg.
	if err := robotgo.SaveJpeg(img, fmt.Sprintf("%s/%s_%s_%s.png", outputDir, number, suffix, randStr)); err != nil {
		return fmt.Errorf("error saving number: %w", err)
	}

	return nil
}
