package driver

import (
	"regexp"
	"strconv"
)

// Battery retrieves the current battery level of the device
// Returns:
//   - An integer representing the battery level (0-100)
//   - If the battery level cannot be retrieved, it returns 0
func (d *driver) Battery() int {
	output, _ := d.Run("dumpsys", "battery", "|", "grep", "level")

	// Regular expression to match the battery level (number)
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(output, -1)

	// Iterate through all matches and convert to integer
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			return num
		}
	}

	// Return 0 if battery level is not found
	return 0
}

// StopCharging disables all charging sources (AC, USB, Wireless)
// This will prevent the device from charging
func (d *driver) StopCharging() {
	d.Run("dumpsys", "battery", "set", "ac", "0")
	d.Run("dumpsys", "battery", "set", "usb", "0")
	d.Run("dumpsys", "battery", "set", "wireless", "0")
}

// StartCharging resets the battery system, enabling charging again
// This restores all charging sources and allows the device to charge
func (d *driver) StartCharging() {
	d.Run("dumpsys", "battery", "reset")
}
