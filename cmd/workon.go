/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/devx-cafe/gh-do/executor"
	"github.com/devx-cafe/gh-do/options"
	"github.com/devx-cafe/gh-do/utils"
	"github.com/spf13/cobra"
)

// Command library for workon
// workonCmd represents the workon command
var workonCmd = &cobra.Command{
	Use:   "workon [issue]",
	Short: "Create or resume a branch to work on an issue",
	Long:  "Creates a new local branch from the remote integration branch. If sucha a branch already exist it will resume work here with a simple checkout.",
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		utils.ValidateGitRepo()

		// First argument must be an integer
		_, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Argument [issue] must be a number")
			fmt.Println(cmd.Usage())
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		workoncmds := map[string]string{
			"remoteOrigin": "git config --get remote.origin.url",
			"gitRepo":      "git rev-parse --show-toplevel",
		}

		if options.Verbose {
			fmt.Println("workon called")
		}

		path, err := executor.RunString(workoncmds["gitRepo"])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(path)

	},
}

func init() {
	rootCmd.AddCommand(workonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
