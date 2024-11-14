package driver

// Input types the specified text at the given coordinates
// Parameters:
//   - x: The x-coordinate to tap
//   - y: The y-coordinate to tap
//   - text: The text to input
func (d *Driver) Input(x, y int, text string) {
	d.Tap(x, y)
	d.Clear(x, y)
	d.Run("am", "broadcast", "-a", "STAR_INPUT_TEXT", "--es", "text", text)
	d.Back()
}

// Clear clears the text at the given coordinates
// Parameters:
//   - x: The x-coordinate to clear
//   - y: The y-coordinate to clear
func (d *Driver) Clear(x, y int) {
	d.Run("am", "broadcast", "-a", "STAR_CLEAR_TEXT")
}
