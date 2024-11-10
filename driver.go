package driver

import "runtime"

// driver represents the core structure for Android UI automation
type driver struct {
	os              string // Operating system name
	shell           string // Shell type (powershell/bash/sh)
	device          string // Connected device ID
	defaultKeyboard string // Default keyboard on device
	deviceInfo      string // Device information string
}

// New creates and initializes a new driver instance
// Returns:
//   - *driver: Configured driver object ready for automation
func New() *driver {
	var d = &driver{
		os: runtime.GOOS,
	}

	// Set shell based on operating system
	if d.os == "windows" {
		d.shell = "powershell" // For Windows 10 and above
	} else if d.os == "android" {
		d.shell = "sh"
	} else {
		d.shell = "bash" // For macOS and Linux
	}

	// Initialize if running on Android
	if d.os == "android" {
		d.initialize()
	}

	return d
}
