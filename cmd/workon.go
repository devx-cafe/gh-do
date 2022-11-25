/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/cli/go-gh"
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

		if options.Verbose {
			fmt.Println("workon called")
		}

		client, _ := gh.RESTClient(nil)
		repo, _ := gh.CurrentRepository()

		response := struct{ Title string }{}
		err := client.Get(fmt.Sprintf("repos/%s/%s/issues/%s", repo.Owner(), repo.Name(), args[0]), &response)
		if err != nil {
			fmt.Println(err)
			return
		}
		name := strings.Join(strings.Fields(strings.ToLower(response.Title)), "-")

		m := regexp.MustCompile("![a-z0-9]")
		mn := regexp.MustCompile("\\.")
		cleanedName := m.ReplaceAllString(name, "")
		cleanedName = mn.ReplaceAllString(cleanedName, "-")

		fmt.Println(fmt.Sprintf("%s-%s", args[0], cleanedName))
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
