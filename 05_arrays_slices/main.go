package main

import "fmt"

func main() {
	// Arrays

	// var fruits [2]string

	// Assign values
	// fruits[0] = "Apple"
	// fruits[1] = "Orange"

	// Declare and assign
	// fruits := [2]string{"Apple", "Orange"}

	// Slices
	fruits := []string{"Apple", "Orange", "Banana", "Grape"}

	fmt.Println(fruits)
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(fruits[2])

	// Several method for array
	fmt.Println(len(fruits)) // array length -> output 4
	fmt.Println(fruits[1:3]) // range -> output: Orange, Banana
}