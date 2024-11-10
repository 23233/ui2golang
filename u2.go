package driver

import (
	"bufio"
	"strings"
)

// Check if the UiAutomator server is running
func (d *driver) checkUiAutomator() (bool, string) {
	output, _ := d.Run("netstat", "-anp", "2>/dev/null")

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ":9008") {
			process := strings.Fields(line)
			if len(process) > 1 {
				pid := strings.Split(process[len(process)-1], "/")[0]
				return true, pid
			}
			break
		}
	}

	return false, ""
}

// Stop the UiAutomator server if it is running
func (d *driver) stopUiAutomator() {
	if running, pid := d.checkUiAutomator(); running {
		d.Run("kill", pid)
	}
}

// Start the UiAutomator server
func (d *driver) startUiAutomator() {
	d.stopUiAutomator()
	go d.Run("CLASSPATH="+U2_PATH, "app_process", "/", "com.wetest.uia2.Main")
}
