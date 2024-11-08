package driver

import "strings"

// StartApp launches an Android application using its package name
// Parameters:
//   - app: full package name of the application (e.g. "com.example.app/.MainActivity")
func (d *driver) StartApp(app string) {
	d.Run("am", "start", "-n", app)
}

// StopApp forcefully stops a running Android application
// Parameters:
//   - app: package name of the application to stop
func (d *driver) StopApp(app string) {
	d.Run("am", "force-stop", app)
}

// RestartApp restarts an Android application by stopping and starting it
// Parameters:
//   - app: full package name of the application to restart
func (d *driver) RestartApp(app string) {
	d.Run("am", "start", "-S", "-n", app)
}

// InstallApp installs an APK file on the Android device
// Parameters:
//   - app: path to the APK file to install
//   - isDel: whether to delete the APK file after installation
func (d *driver) InstallApp(app string, isDel bool) {
	d.Run("pm", "install", app)

	if isDel {
		d.DeleteFile(app)
	}
}

// IsRunning checks if an application is currently running in the foreground
// Parameters:
//   - app: package name of the application to check
// Returns:
//   - bool: true if the app is running, false otherwise
func (d *driver) IsRunning(app string) bool {
	output, err := d.Run("dumpsys", "window", "|", "grep", "-E", "'mCurrentFocus'")
	if err != nil {
		return false
	}

	return strings.Contains(output, app)
}
