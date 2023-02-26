package main

import (
	"fmt"
	"net/http"
)

func main() {
	// "/" is the pathname(url), then function with a response and a pointer to request (to address in memory, where some data is stored).

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err:= fmt.Fprintf(w, "Hello, world!")
		// Fprintf is part of fmt library. It requires ResponseWriter (w)
		//it returns (n int, err error) - number of bytes written and any write error encountered.

		if err != nil {
		fmt.Println(err)
		// We should use n, err.
	}
	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	// Sprintf allows to take different data type and return them as a string.
	})

	_ = http.ListenAndServe(":8080", nil)
	// starts a web-server that listens for requests 
	// localhost:8080 (type this in any browser). First 1024 ports are priviledged, for SUDO only.
	// second argument is a Handler.
}