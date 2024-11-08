package driver

// initialize performs initial setup for the driver:
// - Stores the current keyboard as default
// - Creates temporary directory if it doesn't exist
func (d *driver) initialize() {
	d.defaultKeyboard = d.getCurrentKeyboard()

	if !DirExists(TEMP_PATH) {
		CreateDir(TEMP_PATH)
	}
}
