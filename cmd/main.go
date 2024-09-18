package main

import (
	"git-pilot/internal/runners"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "git-pilot",
		Short: "GitPilot automates your git workflows.",
	}

	git_runner := &runners.GitRunner{}
	var message string
	var createCommitCmd = &cobra.Command{
		Use:   "create-commit",
		Short: "Generate a commit and push it to the remote repository.",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Handle errors more gracefully.
			err := git_runner.StageAndCommit(message)
			if err != nil {
				log.Fatal(err)
			}

			err = git_runner.Push()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	createCommitCmd.Flags().StringVarP(&message, "message", "m", "", "Commit Message.")

	rootCmd.AddCommand(createCommitCmd)
	rootCmd.Execute()
}
