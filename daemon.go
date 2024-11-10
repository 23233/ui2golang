package driver

import (
	"os"

	"github.com/sevlyar/go-daemon"
)

// runDaemon initializes and creates a new daemon context
// Returns:
//   - *daemon.Context: The daemon context if successful, nil if this is the parent process
func runDaemon() (cntxt *daemon.Context) {
	cntxt = &daemon.Context{
		PidFilePerm: 0644,
		LogFilePerm: 0640,
		WorkDir:     ROOT_PATH,
		Umask:       022,
	}

	if f, err := os.OpenFile(DAEMON_PATH, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err == nil {
		f.Close()
		cntxt.LogFileName = DAEMON_PATH
	}

	child, _ := cntxt.Reborn()

	if child != nil {
		return nil
	}

	return cntxt
}

// Daemon runs the provided function as a daemon process
// Parameters:
//   - exec: The function to run in the daemon process
func Daemon(exec func()) {
	cntxt := runDaemon()
	if cntxt == nil {
		return
	}
	defer cntxt.Release()

	exec()

	select {}
}
