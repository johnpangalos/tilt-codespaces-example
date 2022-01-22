package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handle route using handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to new server!")
	})

	fmt.Println("Server started on port 5050.")

	// listen to port
	http.ListenAndServe(":5050", nil)
}
