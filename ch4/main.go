package main

import (
	"fmt"
)

// starting point for chapter 4
func main() {
	fmt.Println("Chapter 4")
	arrays()
}

// 4.1 Arrays
func arrays() {
	fmt.Println("4.1 Arrays")
	// declare and initialize an array of 5 integers
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	// declare an array of 3 strings, initialize 2nd element
	b := [3]string{1: "world"}
	b[0] = "hello"
	b[2] = "!"
	fmt.Println(b)

	//iterate over array
	for i, v := range a {
		fmt.Printf("Element %d is %d\n", i, v)
	}
	//or
	for i := range a {
		fmt.Printf("Element %d is %d\n", i, a[i])
	}
	//or
	for i := 0; i < len(a); i++ {
		fmt.Printf("Element %d is %d\n", i, a[i])
	}
	//multidimensional array
	c := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(c)
	for i, row := range c {
		for j, val := range row {
			fmt.Printf("Element at (%d,%d) is %d\n", i, j, val)
		}
	}
}

// 4.2 Slices
func slices() {
	fmt.Println("4.2 Slices")
}

// 4.3 Maps
func maps() {
	fmt.Println("4.1 Maps")
}
