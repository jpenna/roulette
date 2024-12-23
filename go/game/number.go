package game

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/go-vgo/robotgo"
	"github.com/otiai10/gosseract/v2"
)

var (
	ErrNotANumber    = errors.New("not a number")
	ErrInvalidNumber = errors.New("invalid number")
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

type NumberArea struct {
	bounds image.Rectangle
	points []image.Point
}

func init() {
	client = gosseract.NewClient()
	client.SetWhitelist("0123456789")

	for _, number := range blackNumbers {
		isBlackArr[number] = true
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

	color, err := n.getColor(img)
	if err != nil {
		return 0, fmt.Errorf("error getting color: %w", err)
	}

	return processImage(img, color)
}

func (n *NumberArea) getColor(img image.Image) (Color, error) {
	// Initialize color to first pixel
	var color [3]uint32
	color[0], color[1], color[2], _ = img.At(0, 0).RGBA()

	// Compare with some tolerance as colors might not be exact
	const tolerance uint32 = 0x4000
	var isRed, isBlack, isGreen bool

	if abs(color[0]-0xffff) <= tolerance && color[1] <= tolerance && color[2] <= tolerance {
		isRed = true
	} else if color[0] <= tolerance && abs(color[1]-0xffff) <= tolerance && color[2] <= tolerance {
		isGreen = true
	} else if color[0] <= tolerance && color[1] <= tolerance && color[2] <= tolerance {
		isBlack = true
	}

	// If none found, it's not a number
	if !isRed && !isBlack && !isGreen {
		return 0, ErrNotANumber
	}

	// Check the borders for the same color
	for _, point := range n.points {
		r, g, b, _ := img.At(point.X, point.Y).RGBA()

		// The border color should be the same as the first pixel to be a number
		if color[0] != r || color[1] != g || color[2] != b {
			return 0, ErrNotANumber
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

// Helper function to calculate absolute difference between uint32
func abs(a uint32) uint32 {
	if a < 0x8000 {
		return 0x8000 - a
	}
	return a - 0x8000
}

func processImage(img image.Image, color Color) (int, error) {
	// Only 0 is green
	if color == ColorGreen {
		saveNumber(img, "0")
		return 0, nil
	}

	// Invert image to make the number more readable for OCR
	inverted := invertImage(img)

	client.SetImageFromBytes(robotgo.ToByteImg(inverted))

	text, err := client.Text()
	if err != nil {
		return 0, fmt.Errorf("error performing OCR: %w", err)
	}

	number, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("error parsing number: %w", err)
	}

	if err := saveNumber(img, text); err != nil {
		return 0, fmt.Errorf("error saving number: %w", err)
	}

	// TODO: Remove this
	if err := saveNumber(inverted, text+"_inverted"); err != nil {
		return 0, fmt.Errorf("error saving number: %w", err)
	}

	if err := validateNumber(number, color); err != nil {
		if err := saveNumber(inverted, text+"_invalid"); err != nil {
			log.Printf("error saving number: %v", err)
		}
		return 0, fmt.Errorf("error validating number: %w", err)
	}

	return number, nil
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
		return fmt.Errorf("number should be BLACK but recognized is RED (%d): %w", number, ErrInvalidNumber)
	} else if color == ColorRed && isBlackArr[number] {
		return fmt.Errorf("number should be RED but recognized is BLACK (%d): %w", number, ErrInvalidNumber)
	}

	return nil
}

func saveNumber(img image.Image, number string) error {
	randStr := fmt.Sprintf("%x", rand.Int31())
	file, err := os.Create(fmt.Sprintf("%s_%s.png", number, randStr))
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the image as PNG
	if err := png.Encode(file, img); err != nil {
		return err
	}

	return nil
}
