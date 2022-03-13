package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uin64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// 1. Using var
	// var name string = "Yahfi Ilham" // not best practice
	// var name = "Yahfi Ilham"
	var age int32 = 19
	const isCool = true
	// isCool = false // cannot reasign to const variable

	// 1.1 Shorthand
	// name := "Yahfi Ilham"
	size := 1.3
	// email := "yahfi.ilham@gmail.com"

	name, email := "yahfi Ilham", "yahfi.ilham@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", name) // "%T" represent to type of value
	fmt.Printf("%T\n", age) // for value number by default data type is int, but if we specify type in var declaration its follow it
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size) // by defautl data type is float64
	fmt.Printf("%T\n", email)
}