/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/cli/go-gh"
	"github.com/devx-cafe/gh-do/options"
	"github.com/devx-cafe/gh-do/shell"
	"github.com/devx-cafe/gh-do/utils"
	"github.com/spf13/cobra"
)

// Command library for workon
// workonCmd represents the workon command
<<<<<<< HEAD

var workonCmd = &cobra.Command{
	Use: `workon --new [--title TITLE [--body BODY ]]
	workon ISSUE [--reopen]`,
	Short: "Create or resume a branch to work on an issue",
	Long:  "Creates a new local branch from the remote integration branch. If sucha a branch already exist it will resume work here with a simple checkout.",
	PreRun: func(cmd *cobra.Command, args []string) {

		new := cmd.Flag("new").Changed
		reopen := cmd.Flag("new").Changed
		title := cmd.Flag("new").Changed
		body := cmd.Flag("new").Changed

		if new && reopen {
			fmt.Println("--reopen and --new cannot be used simultaneously")
			os.Exit(0)
		}

=======
// Use: `workon --new [-t,--title TITLE [--b,--body BODY ]]
//
//	workon ISSUE [--reopen]`,
var workonCmd = &cobra.Command{
	Use:   `workon {--new| --reopen}`,
	Short: "Create or resume a branch to work on an issue",
	Long:  "Creates a new local branch from the remote integration branch. If sucha a branch already exist it will resume work here with a simple checkout.",
	PreRun: func(cmd *cobra.Command, args []string) {
>>>>>>> dcc8791 (sync against #19)
		//		utils.ValidateGitRepo()
		//
		//		// First argument must be an integer
		//		_, err := strconv.Atoi(args[0])
		//		if err != nil {
		//			fmt.Fprintln(os.Stderr, "Argument [issue] must be a number")
		//			fmt.Println(cmd.Usage())
		//			os.Exit(1)
		//		}
<<<<<<< HEAD

=======
>>>>>>> dcc8791 (sync against #19)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flags())

		cmd.Flags().Lookup("new")
		if options.Verbose {
			fmt.Println("workon called")
		}

		client, _ := gh.RESTClient(nil)
		repo, _ := gh.CurrentRepository()
		issueID := args[0]

		issueResponse := struct{ Title string }{}
		err := client.Get(fmt.Sprintf("repos/%s/%s/issues/%s", repo.Owner(), repo.Name(), issueID), &issueResponse)
		if err != nil {
			fmt.Println(err)
			return
		}

		branchName := utils.GetBranchName(issueResponse.Title, issueID)

		out, err := shell.RunString(fmt.Sprintf("git checkout -b %s origin/master", branchName))
		if err != nil {
			fmt.Println(out)
			fmt.Println(err.Error())
			os.Exit(0)
		}
		fmt.Println(out)
		shell.RunString(fmt.Sprintf("git checkout -b %s origin/master", branchName))
		out, err := executor.RunString(fmt.Sprintf("git checkout -b %s origin/master", branchName))
		if err != nil {
			fmt.Println(out)
			fmt.Println(err.Error())
			os.Exit(0)
		}
		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(workonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workonCmd.PersistentFlags().String("foo", "", "A help for foo")
	workonCmd.PersistentFlags().BoolVarP(&options.New, "new", "n", false, "Create new issue on GitHub")
	workonCmd.PersistentFlags().BoolVar(&options.ReOpen, "reopen", false, "Reopen closed issue")
	workonCmd.PersistentFlags().StringVarP(&options.Title, "title", "t", "", "Issue title")
	workonCmd.PersistentFlags().StringVarP(&options.Body, "body", "b", "", "Issue body")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
