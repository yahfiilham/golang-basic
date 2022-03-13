package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Golang!</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w , "<h1>About Golang</h1>")
}

func main()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Serving running at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}