package driver

// FileExists checks if a file exists at the given path
// Parameters:
//   - path: Path to check
// Returns:
//   - bool: true if file exists, false otherwise
func (d *driver) FileExists(path string) bool {
	output, err := d.Run("test", "-e", path)
	return err == nil && output == ""
}

// DirExists checks if a directory exists at the given path
// Parameters:
//   - path: Path to check
// Returns:
//   - bool: true if directory exists, false otherwise
func (d *driver) DirExists(path string) bool {
	output, err := d.Run("test", "-d", path)
	return err == nil && output == ""
}

// CreateDir creates a directory and any necessary parent directories
// Parameters:
//   - path: Path where to create directory
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) CreateDir(path string) bool {
	_, err := d.Run("mkdir", "-p", path)
	return err == nil
}

// CreateFile creates a new file with the given text content
// Parameters:
//   - text: Content to write to file
//   - path: Path where to create file
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) CreateFile(text, path string) bool {
	_, err := d.Run("echo", text, ">", path)
	return err == nil
}

// DeleteFile deletes a file or directory recursively
// Parameters:
//   - path: Path to delete
// Returns:
//   - bool: true if successful, false otherwise
func (d *driver) DeleteFile(path string) bool {
	_, err := d.Run("rm", "-rf", path)
	return err == nil
}

// ReadFile reads the content of a file
// Parameters:
//   - path: Path of file to read
// Returns:
//   - string: Content of the file
//   - error: nil if successful, otherwise error details
func (d *driver) ReadFile(path string) (string, error) {
	text, err := d.Run("cat", path)
	if err != nil {
		return "", err
	}
	return text, nil
}
