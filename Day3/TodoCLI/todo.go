package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to CLI Todo App")
	options := []string{}
	options = append(options, "Add", "List", "Mark as Done", "Delete", "Exit")

	for index, option := range options {
		fmt.Println(index, ")", option)
	}

	// fmt.Println(options)
	var input int
	fmt.Print("Choose option:")
	fmt.Scanln(&input)
	fmt.Println("\nYou chose option", input, "\n")

	if input == 1 {

		fmt.Print("Enter task:")
		reader := bufio.NewReader(os.Stdin)
		inputStr, _ := reader.ReadString('\n')
		fmt.Println("\n")
		inputStr = strings.TrimSpace(inputStr)

		result, bool := addTodo(inputStr)
		fmt.Println("Todo added!\n")
		fmt.Println("----Todo List----\n")
		if bool == false {
			fmt.Println("[1]:", result, "(Pending)")
		} else {
			fmt.Println("[1]", result, "(Done)")
		}
	}
}

type TodoList struct {
	Task   string
	Status bool
}

func addTodo(task string) (string, bool) {
	t := new(TodoList)
	t.Task = task
	t.Status = true

	return t.Task, t.Status
}
