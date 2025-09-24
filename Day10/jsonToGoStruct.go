package main

import (
	"encoding/json"
	"fmt"
)

type Todo2 struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}

func main() {
	data := `{"id":1, "task":"Write JSON", "status":true}`

	var todo Todo2
	err := json.Unmarshal([]byte(data), &todo)
	//turn data into byte format and pointer to todo variable so that it stores values of unmarshalled data
	if err != nil {
		panic("Error")
	}

	fmt.Println(todo.ID, todo.Task, todo.Status)
}
