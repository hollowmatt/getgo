package main

import "fmt"

func main() {
	text := "Let's count some words"
	var numSpaces int

	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			numSpaces++
		}
	}

	fmt.Println("found", numSpaces+1, "words")
}
