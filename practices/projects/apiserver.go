package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage")
}

func about(w http.ResponseWriter, r *http.Request) {
	// Set the response content type to HTML
	w.Header().Set("Content-Type", "text/html")

	// Add HTML structure and CSS for centering
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <style>
        body {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            text-align: center;
            font-family: Arial, sans-serif;
        }
        .container {
            text-align: center;
        }
        img {
            margin-top: 20px;
            max-width: 100%;
            height: auto;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Welcome to the About page</h1>
        <p> This is About page of the Website</p>
        <img src='/images/about.png' alt='About Image'>
    </div>
</body>
</html>`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	// Set the response content type to HTML
	w.Header().Set("Content-Type", "text/html")

	// Add HTML structure and CSS for centering
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <style>
        body {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            text-align: center;
            font-family: Arial, sans-serif;
        }
        .container {
            text-align: center;
        }
        img {
            margin-top: 20px;
            max-width: 100%;
            height: auto;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Welcome to the Contact page</h1>
        <p> This is Contact page of the Website</p>
        <img src='/images/contact.png' alt='Contact Image'>
    </div>
</body>
</html>`)
}

func main() {
	fmt.Println("Starting an HTTP server in Go...")

	// Serve static files (like images) from the 'images' directory
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	// Set up routes
	http.HandleFunc("/", homepage)       // Homepage route
	http.HandleFunc("/about", about)     // About route
	http.HandleFunc("/contact", contact) // Contact route

	// Start the server
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}