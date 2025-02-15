package main

import (
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "page/root.html")
	fmt.Println("The user is hitting the root page, '/' ")
}
func about(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "page/about.html")
	fmt.Println("The user is hitting the about page for '/about' ")
}
func contact(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "page/contact.html")
	fmt.Println("The user is hitting the contact page for '/contact' ")
}

func main() {
	fmt.Println("function from the 'net/http' package in Go")
	fmt.Println("handler function in Go is used to process http request")
	http.HandleFunc("/", root)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	fmt.Println("Server started at http://localhost:10000")
	log.Fatal(http.ListenAndServe(":10000", nil))

}
