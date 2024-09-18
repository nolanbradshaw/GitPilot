package runners

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func setupTempDir() (string, error) {
	// Create a temporary directory for testing
	tmpDir, err := ioutil.TempDir("", "gitrunner-test")
	if err != nil {
		return "", err
	}
	return tmpDir, nil
}

// TODO: Simplify test once other git commands are added to the runner.
func TestGitRunnerAdd(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := setupTempDir()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a new GitRunner instance
	gitRunner := &GitRunner{}

	// Create a test file to add to the Git repository
	testFile := filepath.Join(tmpDir, "test.txt")
	err = ioutil.WriteFile(testFile, []byte("Hello, world!"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = tmpDir
	err = cmd.Run()
	if err != nil {
		t.Errorf("git init command failed: %v", err)
	}

	// Run the Add method
	err = gitRunner.Add(".", &tmpDir)
	if err != nil {
		t.Errorf("Add method returned error: %v", err)
	}

	// Verify that the file was added to the Git repository
	cmd = exec.Command("git", "ls-files")
	cmd.Dir = tmpDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("git ls-files command failed: %v", err)
	}

	expectedFile := strings.TrimSpace(string(output))
	if expectedFile != "test.txt" {
		t.Errorf("Expected output %q, got %q", "test.txt", expectedFile)
	}
}
