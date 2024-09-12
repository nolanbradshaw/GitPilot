package main

import (
	"git-pilot/internal/runners"
	"log"
)

func main() {
	git_runner := &runners.GitRunner{}
	err := git_runner.StageAndCommit("Added GitRunner.Status and made StageAndCommit output the status.")
	if err != nil {
		log.Fatal(err)
	}
}
