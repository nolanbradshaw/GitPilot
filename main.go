package main

import (
	"git-pilot/internal/runners"
	"log"
)

func main() {
	git_runner := &runners.GitRunner{}
	err := git_runner.StageAndCommit("Generated using GitPilot.")
	if err != nil {
		log.Fatal(err)
	}

	git_runner.Push()
}
