package executor

import (
	"os/exec"
	"strings"
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
//	reg := []string {"a","b","c"}
//	fmt.Println(strings.Join(strings.Fields(cmdstr), ","))
//	reg := []string cmdargs
	a, b := cmdargs[0], cmdargs[1:]
	cmd := exec.Command(a, b...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
