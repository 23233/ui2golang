package driver

import (
	"bufio"
	"os/exec"
	"strings"
)

// Connect establishes a connection to an Android device
// Parameters:
//   - device: device serial number or identifier
//
// Returns:
//   - error: nil if successful, otherwise:
//   - ErrMultipleDevices if multiple devices are connected
//   - ErrDeviceOffline if the device is offline
//   - Other errors from adb command execution
func (d *driver) Connect(device string) error {
	if d.os == "android" {
		d.device = ""
		return nil
	}

	cmd := exec.Command(d.shell, "adb", "devices")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	out := string(output)
	if strings.Contains(out, "more than one device") {
		return ErrMultipleDevices
	}

	scanner := bufio.NewScanner(strings.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, device) {
			if strings.Contains(line, "offline") {
				return ErrDeviceOffline
			}
			d.device = strings.Fields(line)[0]
			break
		}
	}

	d.initialize()

	return nil
}
