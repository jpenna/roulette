package game

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"time"

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

	debugImgFolder  = "debug"
	resultImgFolder = "result"
	failedImgFolder = "failed"
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

	// Ensure directory exists
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
	if color == ColorBlack {
		// Inverted make the number more readable for OCR (black on white)
		img = invertImage(img)
	}

	if color == ColorRed {
		// Turn it black on white background
		img = blackWhiteImage(img)
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
		return handleFailNumber(img, text, invert, "error parsing number")
	}

	if err := validateNumber(number, color); err != nil {
		return handleFailNumber(img, text, invert, "error validating number")
	}

	if err := saveNumber(img, resultImgFolder, text, invert, true); err != nil {
		fmt.Printf("error saving number: %v", err)
	}

	return number, nil
}

func handleFailNumber(img image.Image, text string, invert bool, errText string) (int, error) {
	if err := saveNumber(img, failedImgFolder, text, invert, false); err != nil {
		log.Printf("error saving number: %v", err)
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

func blackWhiteImage(img image.Image) image.Image {
	bounds := img.Bounds()
	bw := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// Calculate brightness - if close to white (high values), convert to black
			brightness := (r + g + b) / 3
			var pixel uint8
			if brightness > 51400 { // Threshold for "close to white" (200)
				pixel = 0 // Black
			} else {
				pixel = 255 // White
			}

			i := (y-bounds.Min.Y)*bw.Stride + (x-bounds.Min.X)*4
			bw.Pix[i+0] = pixel // R
			bw.Pix[i+1] = pixel // G
			bw.Pix[i+2] = pixel // B
			bw.Pix[i+3] = 255   // A (fully opaque)
		}
	}

	return bw
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

func saveNumber(img image.Image, folder string, number string, invert bool, valid bool) error {
	suffix := ""
	if invert {
		suffix = "_X"
	}
	if !valid {
		suffix += "_I"
	}

	// Ensure directory exists
	var dir = fmt.Sprintf("%s/%s", outputDir, folder)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("error creating output directory: %v", err)
	}

	// randStr := fmt.Sprintf("%x", rand.Int31())
	hash := sha256.New()
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			hash.Write([]byte{byte(r), byte(g), byte(b)})
		}
	}
	randStr := hex.EncodeToString(hash.Sum(nil))[:8]

	filename := fmt.Sprintf("%d_%s%s_%s.png", time.Now().UnixMilli(), number, suffix, randStr)

	img = removeAlpha(img)

	// Saving it to png makes it fully transparent, so using jpeg.
	if err := robotgo.SavePng(img, fmt.Sprintf("%s/%s", dir, filename)); err != nil {
		return fmt.Errorf("error saving number: %w", err)
	}

	return nil
}

func removeAlpha(img image.Image) image.Image {
	if _, _, _, a := img.At(0, 0).RGBA(); a == 255 {
		return img
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	return rgba
}
