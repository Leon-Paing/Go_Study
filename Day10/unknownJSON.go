package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := `{"id":1, "task":"Write some code", "tags":["struct", "json"]}`

	var result map[string]interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		panic("Error converting")
	}

	fmt.Println(result["id"])
	fmt.Println(result["task"])
	fmt.Println(result["tags"])
}
