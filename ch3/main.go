package main

import (
	"fmt"
)

func main() {
	// 3.6 Structs
	type person struct {
		name string
		age  int64
	}

	// How to get an int from the args
	//person.age, _ = strconv.ParseInt(os.Args[2], 10, 64)
	people := []person{
		{
			name: "Dave",
			age:  30,
		},
		{
			name: "Jane",
			age:  25,
		},
	}
	for _, person := range people {
		fmt.Printf("Hello, my name is %s and I am %d years old\n", person.name, person.age)
	}

	// 3.7 Pointers

}
