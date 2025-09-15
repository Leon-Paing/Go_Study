package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to CLI Todo App!\n")
	options := []string{}
	options = append(options, "Add", "List", "Mark as Done", "Delete", "Exit")

	for index, option := range options {
		fmt.Println(index+1, ")", option, "\n")
	}

	// fmt.Println(options)
	var List []TodoList
	var input int
	loop := true
	for loop == true {
		fmt.Print("Choose option:")
		fmt.Scanln(&input)

		switch input {
		case 1:
			fmt.Print("Enter task:")
			reader := bufio.NewReader(os.Stdin)
			inputStr, _ := reader.ReadString('\n')
			fmt.Println("\n")
			inputStr = strings.TrimSpace(inputStr)

			var todo TodoList
			result, status := todo.addTodo(inputStr)
			task := TodoList{Task: result, Status: status}
			List = append(List, task)
			fmt.Println("Todo added!\n")
		case 2:
			fmt.Println("\n----Todo List----\n")
			if len(List) == 0 {
				fmt.Println("No tasks to do! \n")
				break
			} else {
				for i, todo := range List {
					fmt.Println("[", i, "]", todo.Task, "[", statusChecker(todo.Status), "]\n")
				}
			}
		case 5:
			fmt.Println("See you! Good Bye!")
			loop = false

		default:
			fmt.Println("This is default")
		}
	}
	fmt.Println("-------End-------")
}

type TodoList struct {
	Task   string
	Status bool
}

func (t *TodoList) addTodo(task string) (string, bool) {
	t.Task = task
	t.Status = false

	return t.Task, t.Status
}

func statusChecker(stat bool) string {
	if stat == false {
		return "Pending"
	} else {
		return "Done"
	}
}
