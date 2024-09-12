package runners

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

type GitRunner struct{}

func (git_runner *GitRunner) StageAndCommit(message string) error {
	// Stage all changes
	err := git_runner.Add(".")
	if err != nil {
		return err
	}

	log.Print("Staged all changes.")

	status, err := git_runner.Status()
	if err != nil {
		return err
	}

	log.Printf("%s\n", status)

	// TODO: Use git diff as LLM prompt for commit message.
	// git_diff, err := git_runner.Diff()
	// if err != nil {
	// 	return "", err
	// }

	// Commit changes
	_, err = git_runner.Commit(message)
	if err != nil {
		return err
	}

	log.Print("Changes committed.")

	return nil
}

// Stages files in the given path to be commited.
// Returns an error if there is an issue running the command.
// TODO: Should have specific error types for each git error possible?
func (g *GitRunner) Add(path string) error {
	var out bytes.Buffer

	cmd := exec.Command("git", "add", path)
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

// Runs `git commit`.
// Returns the output as a string or an error.
// TODO: Should have specific error types for each git error possible?
func (g *GitRunner) Commit(message string) (string, error) {
	var cmd *exec.Cmd
	if message == "" {
		cmd = exec.Command("git", "commit")
	} else {
		cmd = exec.Command("git", "commit", "-m", message)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		// TODO: Handle errors.
		return "", err
	}

	return out.String(), nil
}

// Runs `git diff`
// Returns the output as a string or an error.
// TODO: Should have specific error types for each git error possible?
func (g *GitRunner) Diff() (string, error) {
	var out bytes.Buffer

	cmd := exec.Command("git", "diff")
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return "", err
	}

	return out.String(), nil
}

// Runs `git status`
// Returns the output as a string or an error.
// TODO: Should have specific error types for each git error possible?
func (git_runner *GitRunner) Status() (string, error) {
	var out bytes.Buffer

	cmd := exec.Command("git", "status")
	cmd.Dir, _ = os.Getwd()
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return "", err
	}

	return out.String(), nil
}
