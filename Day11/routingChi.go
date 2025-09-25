package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos = []Todo{
	{ID: 1, Task: "Learn Chi"},
	{ID: 2, Task: "Write API"},
}

func main() {
	r := chi.NewRouter()

	r.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(todos)
	})

	r.Get("/todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		w.Write([]byte("Todo ID: " + id))
	})

	fmt.Println("Server running at http://localhost:8834")
	http.ListenAndServe(":8834", r)
}
