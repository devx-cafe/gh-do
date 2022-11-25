package utils

import (
	"fmt"
	"os"

	"github.com/devx-cafe/gh-do/executor"
)

// RequiredCurDirRepository ...
// Validates if the command is runnign in a git repository
func ValidateGitRepo() {
	_, err := executor.RunString("git rev-parse --show-toplevel")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Not inside a git repositry ")
		os.Exit(1)
	}
}
