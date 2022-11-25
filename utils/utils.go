package utils

import (
	"fmt"
	"os"

	"github.com/cli/go-gh"
)

var StdCmds = map[string]string{
	"remoteOrigin": "git config --get remote.origin.url",
	"gitRepo":      "git rev-parse --show-toplevel",
	"issueTitle":   "gh issue view %s --json title",
}

// RequiredCurDirRepository ...
// Validates if the command is runnign in a git repository
func ValidateGitRepo() {
	_, err := gh.CurrentRepository()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Not inside a git repositry ")
		os.Exit(1)
	}

}
