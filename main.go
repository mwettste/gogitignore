package main

import (
	"fmt"
	"log"
	"os"
)

var gitignoreContent string = `*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
`

var goModFileName string = "go.mod"

func main() {
	if shouldWriteGitignore() {
		writeGitignore()
	}
}

func shouldWriteGitignore() bool {
	// only write .gitingore if we are in a directory containing a go.mod
	if _, err := os.Stat(goModFileName); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Not writing .gitignore since this does not seem to be a go directory")
			return false
		}
	}

	return true
}

func writeGitignore() {
	if err := os.WriteFile(".gitignore", []byte(gitignoreContent), 0666); err != nil {
		log.Fatalf("Could not write .gitignore %v\n", err)
		return
	}

	fmt.Println("Successfully created .gitignore")
}
