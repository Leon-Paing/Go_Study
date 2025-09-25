package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
		}

		for _, todo := range todos {
			if todo.ID == id {
				json.NewEncoder(w).Encode(todo)
				return
			}
		}

		http.Error(w, "Error", http.StatusNotFound)
	})

	fmt.Println("Server running at http://localhost:8834")
	http.ListenAndServe(":8834", r)
}
