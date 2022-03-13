package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign key value
	// emails["yahfiilham"] = "yahfi.ilham@gmail.com"
	// emails["john"] = "john.doe@gmail.com"
	// emails["alexa"] = "alexa@gmail.com"

	// Declare and add Key Value
	emails := map[string]string{"john": "john@gmail.com", "alexa": "alexa@gmail.com"}
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["john"])
	
	// Delete from map
	delete(emails, "john")
	fmt.Println(emails)
}