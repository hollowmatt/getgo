package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Please provide a filename as argument")
		return
	}

	fileContents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(string(fileContents))
	fmt.Println("found", len(words), "words")
	fmt.Println(words)
}
