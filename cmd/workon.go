/*
Copyright © 2022 Lars & Simon <hey@inc-inc.dk>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/cli/go-gh"
	"github.com/devx-cafe/gh-do/options"
	"github.com/devx-cafe/gh-do/shell"
	"github.com/devx-cafe/gh-do/utils"
	"github.com/spf13/cobra"
)

// Command library for workon
// workonCmd represents the workon command
var workonCmd = &cobra.Command{
	Use: `workon --new [--title TITLE [--body BODY ]]
	workon ISSUE [--reopen]`,
	Short: "Create or resume a branch to work on an issue",
	Long:  "Creates a new local branch from the remote integration branch. If sucha a branch already exist it will resume work here with a simple checkout.",
	PreRun: func(cmd *cobra.Command, args []string) {
		utils.ValidateGitRepo()

		new := cmd.Flag("new").Changed
		reopen := cmd.Flag("reopen").Changed
		title := cmd.Flag("title").Changed
		body := cmd.Flag("body").Changed
		argscount := len(args)

		switch {
		case new && reopen:
			shell.DieGracefully("--reopen and --new cannot be used simultaneously")
		case title && !new:
			shell.DieGracefully("--title can only be used with --new")
		case body && !new:
			shell.DieGracefully("--body can only be used with --new")
		case new && argscount > 0:
			shell.DieGracefully("--new does not take an ISSUE argument")
		case argscount > 0:
			if _, err := strconv.Atoi(args[0]); err != nil {
				shell.DieGracefully("Argument ISSUE must be a number")
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		shell.Vprint("workon called")

		client, _ := gh.RESTClient(nil)
		repo, _ := gh.CurrentRepository()
		issueID := args[0] //TODO

		issueResponse := struct{ Title string }{}
		err := client.Get(fmt.Sprintf("repos/%s/%s/issues/%s", repo.Owner(), repo.Name(), issueID), &issueResponse)
		if err != nil {
			shell.DieGracefully(err)
		}

		branchName := utils.GetBranchName(issueResponse.Title, issueID)

		out, err := shell.RunString(fmt.Sprintf("git checkout -b %s origin/master", branchName))
		if err != nil {
			shell.DieGracefully(out)
		}
	},
}

func init() {
	rootCmd.AddCommand(workonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workonCmd.PersistentFlags().String("foo", "", "A help for foo")
	workonCmd.Flags().BoolVarP(&options.New, "new", "n", false, "Create new issue on GitHub")
	workonCmd.Flags().BoolVar(&options.ReOpen, "reopen", false, "Reopen closed issue")
	workonCmd.Flags().StringVarP(&options.Title, "title", "t", "", "Issue title")
	workonCmd.Flags().StringVarP(&options.Body, "body", "b", "", "Issue body")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
