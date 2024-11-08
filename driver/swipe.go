package driver

import (
	"fmt"
)

type Direction int

const (
	SWIPE_UP Direction = iota
	SWIPE_DOWN
	SWIPE_LEFT
	SWIPE_RIGHT
)

// Swipe performs a swipe gesture across the entire screen
// Parameters:
//   - direction: swipe direction, one of SWIPE_UP/SWIPE_DOWN/SWIPE_LEFT/SWIPE_RIGHT
func (d *driver) Swipe(direction Direction) {
	w, h := d.GetResolution()
	bounds := &Bounds{
		LTX: 0,
		LTY: 0,
		RBX: w,
		RBY: h,
	}

	d.SwipeInRange(bounds, direction, 0, 0.5)
}

// SwipeInRange performs a swipe gesture within a specified boundary
// Parameters:
//   - bounds: boundary coordinates for the swipe area
//   - direction: swipe direction, one of SWIPE_UP/SWIPE_DOWN/SWIPE_LEFT/SWIPE_RIGHT
//   - duration: swipe duration in milliseconds, 0 means using default value 40ms
//   - ratio: swipe distance ratio relative to boundary length, range [0,1]
func (d *driver) SwipeInRange(bounds *Bounds, direction Direction, duration int, ratio float64) {
	if duration == 0 {
		duration = 40 // Default duration if not specified
	}

	width := bounds.RBX - bounds.LTX
	height := bounds.RBY - bounds.LTY

	// Determine start point range based on swipe direction
	var startX, startY int
	switch direction {
	case SWIPE_UP:
		// Start point in bottom half for upward swipe
		startX = GetRandomIntInRange(bounds.LTX+width/4, bounds.LTX+width*3/4)
		startY = GetRandomIntInRange(bounds.LTY+height/2, bounds.RBY-height/4)
	case SWIPE_DOWN:
		// Start point in top half for downward swipe
		startX = GetRandomIntInRange(bounds.LTX+width/4, bounds.LTX+width*3/4)
		startY = GetRandomIntInRange(bounds.LTY+height/4, bounds.LTY+height/2)
	case SWIPE_LEFT:
		// Start point in right half for leftward swipe
		startX = GetRandomIntInRange(bounds.LTX+width/2, bounds.RBX-width/4)
		startY = GetRandomIntInRange(bounds.LTY+height/4, bounds.LTY+height*3/4)
	case SWIPE_RIGHT:
		// Start point in left half for rightward swipe
		startX = GetRandomIntInRange(bounds.LTX+width/4, bounds.LTX+width/2)
		startY = GetRandomIntInRange(bounds.LTY+height/4, bounds.LTY+height*3/4)
	}

	// Calculate swipe distance based on rectangle dimensions
	var swipeDistance int
	switch direction {
	case SWIPE_UP, SWIPE_DOWN:
		swipeDistance = int(float64(height) * ratio)
	case SWIPE_LEFT, SWIPE_RIGHT:
		swipeDistance = int(float64(width) * ratio)
	}

	// Calculate end point based on direction
	var endX, endY int
	switch direction {
	case SWIPE_UP:
		endX = startX + GetRandomIntInRange(-20, 20)
		endY = startY - swipeDistance
		// Ensure end point doesn't exceed boundary
		if endY < bounds.LTY {
			endY = bounds.LTY
		}
	case SWIPE_DOWN:
		endX = startX + GetRandomIntInRange(-20, 20)
		endY = startY + swipeDistance
		if endY > bounds.RBY {
			endY = bounds.RBY
		}
	case SWIPE_LEFT:
		endX = startX - swipeDistance
		endY = startY + GetRandomIntInRange(-20, 20)
		if endX < bounds.LTX {
			endX = bounds.LTX
		}
	case SWIPE_RIGHT:
		endX = startX + swipeDistance
		endY = startY + GetRandomIntInRange(-20, 20)
		if endX > bounds.RBX {
			endX = bounds.RBX
		}
	}

	// Execute swipe command
	c := fmt.Sprintf("%d %d %d %d %d", startX, startY, endX, endY, duration)
	d.Run("input", "swipe", c)
}
