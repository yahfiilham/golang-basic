package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int

	firstName, lastName, city, gender string
	age int
}

// Greeting method (value receiver)
func (person Person) greet() string {
	return "Hello my name is" + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age) + " years old."
}

// has birthday method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

func main() {
	// Init person using struct
	person1 := Person{ firstName: "Yahfi", lastName: "Ilham", city: "Sleman", gender: "Male", age: 19}

	// Alternative
	// person1 := Person{ "Yahfi", "Ilham", "Sleman", "Male", 19}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)

	// person1.age += 5
	// fmt.Println(person1)

	person1.hasBirthday()
	fmt.Println(person1.greet())
}