package runners

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type GitRunner struct{}

func (g *GitRunner) RunGitDiff() (string, error) {
	cmd := exec.Command("git", "diff")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func (g *GitRunner) Commit(message string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", message)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// No changes were added.
	if strings.Contains(out.String(), "nothing to commit, working tree clean") {
		log.Println("No changes were added, nothing to commit.")
		return "", fmt.Errorf("nothing to commit, working tree clean")
	}

	// TODO: Check that the commit didn't send some other type of error.

	return out.String(), nil
}
