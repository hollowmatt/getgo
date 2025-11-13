package main

import (
	"fmt"
	"sort"
)

// starting point for chapter 4
func main() {
	fmt.Println("Chapter 4")
	arrays(false)
	slices(true)
	maps(false)
}

// 4.1 Arrays
// Arrays are fixed size collections of elements of the same type
func arrays(printElements bool) {
	if printElements {
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
}

// 4.2 Slices
// Slices are dynamic, flexible views into the elements of an array (and are ordered)
// Use slices to make lists
func slices(printElements bool) {
	if printElements {
		fmt.Println("4.2 Slices")
		//declare and initialize a slice of strings with literals
		//looks just like an array but without the size
		s := []string{"apple", "banana", "cherry"}
		fmt.Println(s)
		//now an integer slice
		sliceInt := []int{10, 30, 50, 40, 20}
		sliceInt = append(sliceInt, 60) //add an element
		//sort it
		sort.Slice(sliceInt, func(i, j int) bool {
			return sliceInt[i] < sliceInt[j]
		})
		fmt.Println(sliceInt)
		//declaring with make
		madeSlice := make([]float64, 5) //length 5
		fmt.Println(madeSlice)

	}
}

// 4.3 Maps
// Maps are unordered collections of key-value pairs - like dictionaries in other languages
// use maps to make sets
func maps(printElements bool) {
	if printElements {
		fmt.Println("4.1 Maps")
	}
}
