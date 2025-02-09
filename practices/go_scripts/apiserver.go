package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage")
	fmt.Println("When you access the server locally on port 10000, you should see (Example, http://localhost:10000): homepage")
}

func main() {
	fmt.Println("Setting up an API server...")
	http.HandleFunc("/", homepage)
	err := http.ListenAndServe("localhost:10000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}


## Useful GO commands,
// go doc http 
// go doc http ListenAndServe
// go doc http DefaultServeMux
