package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "words.txt"
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(string(fileContents))
	fmt.Println("found", len(words), "words")
	fmt.Println(words)
}
