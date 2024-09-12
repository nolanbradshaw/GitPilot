package runners

import (
	"bytes"
	"os/exec"
)

type GitRunner struct{}

// Stages files in the given path to be commited.
// Returns an error if there is an issue running the command.
// TODO: Should have specific error types for each git error possible?
func (g *GitRunner) Add(path *string) error {
	var out bytes.Buffer
	cmd := exec.Command("git", "add", *path)
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
	cmd := exec.Command("git", "diff")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
