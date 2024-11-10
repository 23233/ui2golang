package driver

import "strconv"

// Tap performs a tap action at the specified coordinates.
// Parameters:
//   - x: The x-coordinate to tap.
//   - y: The y-coordinate to tap.
func (d *driver) Tap(x, y int) {
	px := strconv.Itoa(x)
	py := strconv.Itoa(y)
	d.Run("input", "tap", px, py)
}

// LongTap performs a long tap action at the specified coordinates.
// Parameters:
//   - x: The x-coordinate to long tap.
//   - y: The y-coordinate to long tap.
func (d *driver) LongTap(x, y int) {
	px := strconv.Itoa(x)
	py := strconv.Itoa(y)
	d.Run("input", "swipe", px, py, px, py, "800")
}
