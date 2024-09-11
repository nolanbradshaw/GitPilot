package main

import (
	"fmt"
	"git-pilot/internal/runners"
	"log"
)

func main() {
	git_runner := &runners.GitRunner{}
	diff, err := git_runner.RunGitDiff()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(diff)
}
