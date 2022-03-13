package main

import (
	"fmt"
)

func main() {
	x := 15
	y := 15

	if x <= y {
		fmt.Printf("%d is less than or equal to %d \n", x, y)
	} else {
		fmt.Printf("%d is less than or equal to %d \n", y, x)
	}

	// else if	

	color := "red"
	// color = "blue"
	// color = "greens"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT red or blue")
	}

	// Switch

	animal := "cat"
	// animal = "dog"
	animal = "monkey"

	switch animal {
	case "cat":
		fmt.Println("animal is cat")
	case "dog":
		fmt.Println("animal is dog")
	default:
		fmt.Println("animal is NOT cat or dog")
	}
}