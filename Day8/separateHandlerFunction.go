package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", HelloHandler)

	fmt.Println("Server running on http://localhost:8834")
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from different function")
}
