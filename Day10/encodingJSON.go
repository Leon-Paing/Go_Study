package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}

func main() {
	todo := Todo{ID: 1, Task: "Learn Go", Status: false}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

	pretty, _ := json.MarshalIndent(todo, "", "  ")
	fmt.Println(string(pretty))

}
