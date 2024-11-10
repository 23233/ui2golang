package driver

import (
	"strings"
	"time"
)

// StartApp launches an Android application using its package name
// Parameters:
//   - app: full package name of the application (e.g. "com.example.app/.MainActivity")
//
// Returns:
//   - bool: true if app successfully started within timeout, false otherwise
func (d *driver) StartApp(app string) bool {
	d.StopApp(app)

	activity := d.getMainActivity(app)
	d.Run("am", "start", "-n", activity)

	for i := 0; i < WAIT_TIMEOUT; i++ {
		if d.isRunning(app) {
			return true
		}
		time.Sleep(time.Second)
	}
	return false
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
	d.StopApp(app)
	d.StartApp(app)
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

// UninstallApp uninstalls an installed Android application
// Parameters:
//   - app: package name of the application to uninstall
func (d *driver) UninstallApp(app string) {
	d.Run("pm", "uninstall", app)
}

// IsRunning checks if an application is currently running in the foreground
// Parameters:
//   - app: package name of the application to check
//
// Returns:
//   - bool: true if the app is running, false otherwise
func (d *driver) isRunning(app string) bool {
	output, err := d.Run("dumpsys", "window", "|", "grep", "-E", "'mCurrentFocus'")
	if err != nil {
		return false
	}

	return strings.Contains(output, app)
}

// getMainActivity returns the main activity of an Android application
// Parameters:
//   - app: full package name of the application (e.g. "com.example.app/.MainActivity")
//
// Returns:
//   - string: main activity of the application
func (d *driver) getMainActivity(app string) string {
	output, _ := d.Run("cmd", "package", "resolve-activity", "--brief", app)
	activity := strings.Split(output, "\n")[1]

	return activity
}
