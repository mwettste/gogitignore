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

func main() {
	if err := os.WriteFile(".gitignore", []byte(gitignoreContent), 0666); err != nil {
		log.Fatalf("Could not write .gitignore %v\n", err)
		return
	}

	fmt.Println("Successfully created .gitignore")
}
