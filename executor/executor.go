package executor

import (
	"os/exec"
)

// Run a terminal command
// First argument is the command, the second are flags
// Run("git","pull")
func Run(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func RunString(str string) (string, error) {
	
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
