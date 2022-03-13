package main

import "fmt"

func main() {
	ids := []int{12, 34, 45, 56, 67}

	// Loop through ids 
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// NO using index 
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("SUM", sum)

	// Range with map
	emails := map[string]string{"john": "john@gmail.com", "alexa": "alexa@gmail.com"}

	for key, val := range emails {
		fmt.Printf("%s: %s\n", key, val)
	}

	for key := range emails {
		fmt.Println("Name:", key)
	}

}