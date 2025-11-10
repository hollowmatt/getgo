package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Please provide a filename as argument")
		return
	}

	//update to allow multiple files
	for _, filename := range os.Args[1:] {
		fileContents, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(fileContents)
		scanner.Split(bufio.ScanWords)
		var wordCount int

		for scanner.Scan() {
			wordCount++
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fileContents.Close()
		fmt.Printf("%s found %d words \n", filename, wordCount)
	}
}
