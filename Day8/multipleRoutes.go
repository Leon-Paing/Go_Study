package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)

	fmt.Println("Server running at http://localhost:8834")
	http.ListenAndServe(":8834", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Home!")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is About Page!")
}
