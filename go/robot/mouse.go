package robot

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

// Point represents a 2D coordinate
type Point struct {
	X, Y float64
}

func MoveTo(x, y int) {
	moveBezier(x, y)
}

// moveBezier moves the mouse in a natural-looking curve using cubic Bézier interpolation
func moveBezier(destX, destY int) {
	start := getCurrentMousePos()
	end := Point{float64(destX), float64(destY)}

	// Generate control points for the Bézier curve
	dist, ctrl1, ctrl2 := generateControlPoints(start, end)

	// Number of steps for the movement (adjust for speed)
	steps := int(dist) * rand.Intn(400) / 1000

	// Move the mouse along the Bézier curve
	count := 0 // Use count so it will end in the correct destination
	for t := 0.0; count <= steps; t += 1.0 / float64(steps) {
		x, y := cubicBezier(start, ctrl1, ctrl2, end, t)
		robotgo.Move(int(x), int(y))

		// Add small random delay for more natural movement
		delay := time.Duration(rand.Float64()*2+1) * time.Millisecond
		time.Sleep(delay)

		count++
	}
}

// getCurrentMousePos gets the current mouse position as a Point
func getCurrentMousePos() Point {
	x, y := robotgo.Location()
	return Point{float64(x), float64(y)}
}

// generateControlPoints creates two control points for the Bézier curve
func generateControlPoints(start, end Point) (float64, Point, Point) {
	// Calculate distance between start and end
	dx := end.X - start.X
	dy := end.Y - start.Y
	dist := math.Sqrt(dx*dx + dy*dy)

	// Add some randomness to control points
	randFactor := dist * 0.4 // Adjust this value to control curve intensity

	// Create control points at roughly 1/3 and 2/3 of the path with some randomness
	ctrl1 := Point{
		X: start.X + dx*0.3 + (rand.Float64()-0.5)*randFactor,
		Y: start.Y + dy*0.3 + (rand.Float64()-0.5)*randFactor,
	}

	ctrl2 := Point{
		X: start.X + dx*0.7 + (rand.Float64()-0.5)*randFactor,
		Y: start.Y + dy*0.7 + (rand.Float64()-0.5)*randFactor,
	}

	return dist, ctrl1, ctrl2
}

// cubicBezier calculates a point along a cubic Bézier curve at time t
func cubicBezier(start, ctrl1, ctrl2, end Point, t float64) (float64, float64) {
	// Cubic Bézier formula
	mt := 1 - t
	mt2 := mt * mt
	mt3 := mt2 * mt
	t2 := t * t
	t3 := t2 * t

	x := mt3*start.X + 3*mt2*t*ctrl1.X + 3*mt*t2*ctrl2.X + t3*end.X
	y := mt3*start.Y + 3*mt2*t*ctrl1.Y + 3*mt*t2*ctrl2.Y + t3*end.Y

	return x, y
}
