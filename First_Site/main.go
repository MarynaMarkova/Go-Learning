package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request){
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

// Divide divides one value into another and returns message with result
func Divide(w http.ResponseWriter, r *http.Request){
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

// divideValues divides one float by another and returs the result of division
func divideValues(x, y float32) (float32, error) {
	if y <=0 {
	err := errors.New("Cannot divide by zero")
	return 0, err
	}
	
	result := x / y
	return result, nil
}
// addValues adds two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}