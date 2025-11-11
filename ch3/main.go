package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("Please provide a name and age as arguments")
		return
	}

	var person struct {
		name string
		age  int64
	}
	person.name = os.Args[1]
	person.age, _ = strconv.ParseInt(os.Args[2], 10, 64)

	fmt.Printf("Hello, my name is %s and I am %d years old.", person.name, person.age)

}
