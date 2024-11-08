package driver

// Input types the specified text at the given coordinates
// Parameters:
//   - x: The x-coordinate to tap
//   - y: The y-coordinate to tap
//   - text: The text to input
func (d *driver) Input(x, y int, text string) {
	d.SwitchAdbKeyboard()
	d.Tap(x, y)
	d.Clear(x, y)
	// time.Sleep(time.Millisecond * 500)
	d.Run("am", "broadcast", "-a", "STAR_INPUT_TEXT", "--es", "msg", text)
	// time.Sleep(time.Millisecond * 500)
	// d.Back()
	d.SwitchDefaultKeyboard()
}

// Clear clears the text at the given coordinates
// Parameters:
//   - x: The x-coordinate to clear
//   - y: The y-coordinate to clear
func (d *driver) Clear(x, y int) {
	d.Run("am", "broadcast", "-a", "STAR_CLEAR_TEXT")
}
