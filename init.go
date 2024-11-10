package driver

import "strings"

// initialize performs initial setup for the driver:
//  - Downloading and installing UiAutomator service if needed
//  - Installing and enabling ADB keyboard if needed
//  - Starting UiAutomator service
//  - Storing current keyboard as default
//  - Switching to ADB keyboard
//  - Creating temp directory if needed
func (d *driver) initialize() {
	if !d.FileExists(U2_PATH) {
		d.DownloadFile(U2_URL, U2_PATH)
	}

	imeList, _ := d.Run("ime", "list", "-s")
	if !strings.Contains(imeList, ADB_KEYBOARD) {
		d.DownloadFile(ADB_KEYBOARD_URL, ROOT_PATH+"/star-ime.apk")
		d.InstallApp(ROOT_PATH+"/star-ime.apk", true)
		d.Run("ime", "enable", ADB_KEYBOARD)
	}

	d.startUiAutomator()

	d.defaultKeyboard = d.getCurrentKeyboard()

	d.SwitchAdbKeyboard()

	if d.os != "android" && !DirExists(TEMP_PATH) {
		CreateDir(TEMP_PATH)
	}
}

// Cleanup performs cleanup after the driver:
//  - Stopping UiAutomator service
//  - Restoring default keyboard
func (d *driver) Cleanup() {
	d.stopUiAutomator()

	d.SwitchDefaultKeyboard()
}
