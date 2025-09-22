// http.HandleFunc("/", handler) : register a handler for /.

// http.ResponseWriter : lets you write a response back to the client.

// *http.Request : contains all details about the incoming request.

// http.ListenAndServe(":8080", nil) : starts the web server.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Go web!")
	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
