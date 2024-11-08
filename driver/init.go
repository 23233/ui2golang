package driver

// Initializes
func (d *driver) initialize() {
	d.defaultKeyboard = d.getCurrentKeyboard()
}
