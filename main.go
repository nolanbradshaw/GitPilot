package main

import (
	"git-pilot/internal/runners"
	"log"
)

func main() {
	git_runner := &runners.GitRunner{}
	err := git_runner.StageAndCommit("Added StageAndCommit func to stage changes and run git commit.")
	if err != nil {
		log.Fatal(err)
	}
}
