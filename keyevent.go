package driver

import "fmt"

// KeyEvent sends a key event with the specified keycode
// Parameters:
//   - keyCode: The Android key code to send
// Returns:
//   - bool: true if successful, false otherwise
func (d *Driver) KeyEvent(keyCode KeyCode) bool {
	if output, err := d.Run("input", "keyevent", fmt.Sprintf("%d", keyCode)); err != nil || output != "" {
		return false
	}
	return true
}

// Home simulates pressing the home button
// Returns:
//   - bool: true if successful, false otherwise
func (d *Driver) Home() bool {
	return d.KeyEvent(KEYCODE_HOME)
}

// Back simulates pressing the back button
// Returns:
//   - bool: true if successful, false otherwise
func (d *Driver) Back() bool {
	return d.KeyEvent(KEYCODE_BACK)
}

// Enter simulates pressing the enter key
// Returns:
//   - bool: true if successful, false otherwise
func (d *Driver) Enter() bool {
	return d.KeyEvent(KEYCODE_ENTER)
}

// Search simulates pressing the search button
// Returns:
//   - bool: true if successful, false otherwise
func (d *Driver) Search() bool {
	return d.KeyEvent(KEYCODE_SEARCH)
}

// Menu simulates pressing the menu button
// Returns:
//   - bool: true if successful, false otherwise
func (d *Driver) Menu() bool {
	return d.KeyEvent(KEYCODE_MENU)
}

// VolumeUp simulates pressing the volume up button
// Returns:
//   - bool: true if successful, false otherwise
func (d *Driver) VolumeUp() bool {
	return d.KeyEvent(KEYCODE_VOLUME_UP)
}

// VolumeDown simulates pressing the volume down button
// Returns:
//   - bool: true if successful, false otherwise
func (d *Driver) VolumeDown() bool {
	return d.KeyEvent(KEYCODE_VOLUME_DOWN)
}

// Power simulates pressing the power button
// Returns:
//   - bool: true if successful, false otherwise
func (d *Driver) Power() bool {
	return d.KeyEvent(KEYCODE_POWER)
}

// Reboot simulates a device reboot command
// This will restart the entire device
func (d *Driver) Reboot() {
	d.Run("reboot")
}

// PowerOff simulates powering off the device
// This will shut down the entire device
func (d *Driver) PowerOff() {
	d.Run("poweroff")
}
