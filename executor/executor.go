package executor

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/devx-cafe/gh-do/options"
)

// Run a terminal command
// First argument is the command, the second are flags
// Run("git","pull")
func Run(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func RunString(cmdstr string) (string, error) {
	cmdargs := strings.Fields(cmdstr)
	a, b := cmdargs[0], cmdargs[1:]

	if options.Verbose {
		fmt.Println(cmdstr)
	}
	cmd := exec.Command(a, b...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
