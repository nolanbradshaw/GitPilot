package main

import (
	"git-pilot/internal/runners"
	"log"
)

func main() {
	git_runner := &runners.GitRunner{}
	err := git_runner.StageAndCommit("Git status summary")
	if err != nil {
		log.Fatal(err)
	}
}
