package driver

// getCurrentKeyboard retrieves the current keyboard input method
func (d *Driver) getCurrentKeyboard() string {
	ime, _ := d.Run("settings", "get", "secure", "default_input_method")
	return ime
}

// SwitchKeyBoard switches the keyboard input method to the specified IME
// Parameters:
//   - ime: The input method to switch to
// Returns:
//   - bool: true if the switch was successful, false otherwise
func (d *Driver) switchKeyboard(ime string) bool {
	_, err := d.Run("ime", "set", ime)
	return err == nil
}

// SwitchAdbKeyboard switches the keyboard input method to the ADB keyboard
// Returns:
//   - bool: true if the switch was successful, false otherwise
func (d *Driver) SwitchAdbKeyboard() bool {
	return d.switchKeyboard(ADB_KEYBOARD)
}

// SwitchDefaultKeyboard switches the keyboard input method to the default keyboard
// Returns:
//   - bool: true if the switch was successful, false otherwise
func (d *Driver) SwitchDefaultKeyboard() bool {
	return d.switchKeyboard(d.defaultKeyboard)
}
