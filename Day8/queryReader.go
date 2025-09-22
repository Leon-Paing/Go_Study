package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", QueryReader)

	fmt.Println("Server running at port 8834...")
	http.ListenAndServe(":8834", nil)
}

func QueryReader(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Greetings, %s!", name)
}
