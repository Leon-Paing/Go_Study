package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo3 struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

func main() {
	file, _ := os.Create("todo.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	todo := Todo3{ID: 2, Task: "Write and read file"}
	encoder.Encode(todo)

	file2, _ := os.Open("todo.json")
	defer file2.Close()

	var readTodo Todo3
	json.NewDecoder(file2).Decode(&readTodo)
	fmt.Println(readTodo)
}
