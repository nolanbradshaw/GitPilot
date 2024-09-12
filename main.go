package main

import (
	"git-pilot/internal/runners"
	"log"
)

func main() {
	git_runner := &runners.GitRunner{}
	err := git_runner.StageAndCommit("Run add first")
	if err != nil {
		log.Fatal(err)
	}
}
