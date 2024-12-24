package utils

import "image"

// ProcessBlack inverts the colors to make the number black on white background
func ProcessBlack(img image.Image) image.Image {
	bounds := img.Bounds()
	inverted := image.NewRGBA(bounds)

	// Process the original image content with inversion
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

// ProcessRed adds a margin to the image, make it all black and white, then downscales it (worked better).
func ProcessRed(img image.Image) image.Image {
	bounds := img.Bounds()

	// Add margin
	bigger, marginX, marginY := biggerSize(img)

	// Process the original image content
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

			i := (y-bounds.Min.Y+marginY)*bigger.Stride + (x-bounds.Min.X+marginX)*4
			bigger.Pix[i+0] = pixel // R
			bigger.Pix[i+1] = pixel // G
			bigger.Pix[i+2] = pixel // B
			bigger.Pix[i+3] = 255   // A (fully opaque)
		}
	}

	// Downscale the image
	return downscaleImage(bigger)
}

func biggerSize(img image.Image) (*image.RGBA, int, int) {
	bounds := img.Bounds()
	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()

	// Calculate new dimensions with 30% margin
	marginX := int(float64(originalWidth) * 0.3)
	marginY := int(float64(originalHeight) * 0.3)
	newWidth := originalWidth + 2*marginX
	newHeight := originalHeight + 2*marginY

	bigger := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// Fill entire image with white first
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			i := y*bigger.Stride + x*4
			bigger.Pix[i+0] = 255 // R
			bigger.Pix[i+1] = 255 // G
			bigger.Pix[i+2] = 255 // B
			bigger.Pix[i+3] = 255 // A
		}
	}

	return bigger, marginX, marginY
}

func downscaleImage(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()

	// Calculate new dimensions (50% smaller)
	newWidth := int(float64(originalWidth) * 0.5)
	newHeight := int(float64(originalHeight) * 0.5)

	// Create new image with decreased resolution
	downscaled := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// Fill with white background first
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			i := y*downscaled.Stride + x*4
			downscaled.Pix[i+0] = 255 // R
			downscaled.Pix[i+1] = 255 // G
			downscaled.Pix[i+2] = 255 // B
			downscaled.Pix[i+3] = 255 // A
		}
	}

	// Scale the image using bilinear interpolation
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			// Map new coordinates back to original image
			srcX := float64(x) * float64(originalWidth-1) / float64(newWidth-1)
			srcY := float64(y) * float64(originalHeight-1) / float64(newHeight-1)

			// Get integer parts
			x0 := int(srcX)
			y0 := int(srcY)
			x1 := x0 + 1
			y1 := y0 + 1

			// Ensure we don't go out of bounds
			if x1 >= originalWidth {
				x1 = originalWidth - 1
			}
			if y1 >= originalHeight {
				y1 = originalHeight - 1
			}

			// Get fractional parts
			fx := srcX - float64(x0)
			fy := srcY - float64(y0)

			// Get colors at four corners
			c00 := img.At(x0, y0)
			c10 := img.At(x1, y0)
			c01 := img.At(x0, y1)
			c11 := img.At(x1, y1)

			// Convert to RGBA values
			r00, g00, b00, a00 := c00.RGBA()
			r10, g10, b10, a10 := c10.RGBA()
			r01, g01, b01, a01 := c01.RGBA()
			r11, g11, b11, a11 := c11.RGBA()

			// Bilinear interpolation for each channel
			r := uint8((float64(r00>>8)*(1-fx)*(1-fy) + float64(r10>>8)*fx*(1-fy) +
				float64(r01>>8)*(1-fx)*fy + float64(r11>>8)*fx*fy))
			g := uint8((float64(g00>>8)*(1-fx)*(1-fy) + float64(g10>>8)*fx*(1-fy) +
				float64(g01>>8)*(1-fx)*fy + float64(g11>>8)*fx*fy))
			b := uint8((float64(b00>>8)*(1-fx)*(1-fy) + float64(b10>>8)*fx*(1-fy) +
				float64(b01>>8)*(1-fx)*fy + float64(b11>>8)*fx*fy))
			a := uint8((float64(a00>>8)*(1-fx)*(1-fy) + float64(a10>>8)*fx*(1-fy) +
				float64(a01>>8)*(1-fx)*fy + float64(a11>>8)*fx*fy))

			// Set pixel in downscaled image
			i := y*downscaled.Stride + x*4
			downscaled.Pix[i+0] = r
			downscaled.Pix[i+1] = g
			downscaled.Pix[i+2] = b
			downscaled.Pix[i+3] = a
		}
	}

	return downscaled
}
