package game

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func IsReadyToBet(x, y int) (bool, error) {
	img, err := robotgo.CaptureImg(x, y, 1, 1)
	if err != nil {
		return false, fmt.Errorf("error capturing image: %w", err)
	}

	pixel := img.At(0, 0)
	r, g, b, _ := pixel.RGBA()

	greenMargin := g - 12850 // 50 is the margin of error

	// Check if green component is significantly higher than red and blue
	return greenMargin > r && greenMargin > b, nil
}
