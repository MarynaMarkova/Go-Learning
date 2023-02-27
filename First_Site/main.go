package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r  *http.Request){
	fmt.Fprintf(w, "This is the home page")
}
//functions with capital letters it's visible outside the package (public analogue), with small letter - only inside the package (private analogue).

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
	// _ before function ignores an error (_, _ - ignores two)
	// %d - is a placeholder for int
}

// addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}