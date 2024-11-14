package driver

import (
	"fmt"
	"os/exec"
	"strings"
)

// pcOnlyCommands is a map that defines commands only executable on a PC.
var pcOnlyCommands = map[string]string{
	"push":      "push",
	"pull":      "pull",
	"install":   "install",
	"uninstall": "uninstall",
	"reboot":    "reboot",
	"forward":   "forward",
	"reverse":   "reverse",
	"backup":    "backup",
	"restore":   "restore",
}

// Run executes an adb command.
// Parameters:
//   - cmd: The command to execute.
//   - args: Additional arguments for the command.
//
// Returns:
//   - string: The output of the command.
//   - error: An error object if the command execution fails.
func (d *Driver) Run(cmd string, args ...string) (string, error) {
	var argv []string

	if d.os != "android" {
		if _, exists := pcOnlyCommands[cmd]; exists {
			argv = append(argv, []string{"adb", "-s", d.device, cmd}...)
			argv = append(argv, args...)
		} else {
			newArgs := fmt.Sprintf(`"%s %s"`, cmd, strings.Join(args, " "))
			argv = append(argv, []string{"adb", "-s", d.device, "shell", newArgs}...)
		}
	} else {
		cmd = strings.Join(append([]string{cmd}, args...), " ")
		argv = append(argv, "-c", cmd)
	}

	command := exec.Command(d.shell, argv...)
	output, err := command.CombinedOutput()

	if err != nil {
		return string(output), err
	}

	return strings.TrimSpace(string(output)), nil
}
