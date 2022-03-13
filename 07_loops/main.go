package main

import "fmt"

func main() {
	// Long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)

	// 	i++
	// }

	// Short method
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Number", i)
	// }

	// Fizzbuzz
	for i:= 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}