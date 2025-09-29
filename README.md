# 🦁 Learn Go Programming with Me

### This is the Go Programming guide for **beginners to experts**.

> 💡 Remember: _Consistency is Key!_

<p align="center">
	<img src="https://media2.giphy.com/media/v1.Y2lkPTc5MGI3NjExaHo3OTg3b3Qya2F5NGFjY3lmODcya3lncDRraW1qOTZmbmRmMmMzMyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/PhTSmzCqkliqIJ9ZtZ/giphy.gif">
</p>

---

## 📚 Table of Contents

[Day 1](#day-1) | Introduction to Go & Setup <br>
[Day 2](#day-2) | Dealing with data <br>
[Day 3](#day-3) | Custom algorithms <br>
[Day 4](#day-4) | Go routines and Channels <br>
[Day 5](#day-5) | Context <br>
[Day 6](#day-6) | Testing <br>
[Day 7](#day-7) | Manual testing for CLITodo App <br>
[Day 8](#day-8) | Introduction to "net/http" <br>
[Day 9](#day-9) | Middleware <br>
[Day 10](#day-10) | "encoding/json" <br>
[Day 11](#day-11) | Routing with Chi/Echo <br>
[Day 12](#day-12) | Simple POST/GET Todo API <br>
[Day 13](#day-13) | Middleware and Validation with echo <br>
[Day 14](#day-14) | CRUD Todo API (In-memory) <br>

---

## Day 1

### Introduction

- **Topics Covered:**
  - Installing Go
  - Learn go run, go build, go mod init
  - Review variables, types, loops, functions
  - Writing your first Go program

_Example Code_

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

## Day 2

### Dealing with data

- **Topics Covered:**
  - Study structs, slices, maps, methods, interfaces
  - Mini project named car-manager

_Example Code_

```go
type Car struct {
	Brand string
	Year  int
	Price float64
}

func (c Car) Detail() (string, int, float64) {
	return c.Brand, c.Year, c.Price
}

func (c Car) Discount(percent float64) float64 {
	return c.Price - ((percent * 100) / c.Price)
}

func (c Car) Value() float64 {
	return c.Price
}

type Valuable interface {
	Value() float64
}
```

## Day 3

### Todo List CLI App

- **Topics Covered:**
  - Implement custom made stack and reverse string
  - Mini project named TodoCLI

_Example Code_

```go
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

func (t *TodoList) markDone() {
	t.Status = true
}
```

## Day 4

### Go routines and Channels

- **Topics Covered:**
  - Understand about go routines
  - Learn why channels are important
  - Know difference between buffered and unbuffered channels
  - Channel sequence
  - Sync and WaitGroup

_Example Code_

```go
for i := 0; i <= 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Message:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Message:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout!..Finished")
		}
	}
```

## Day 5

### Learn about Context

- **Topics Covered:**
  - Context with Timeout
  - Context with Cancel
  - Context with Deadline
  - Context with Value

_Example Code_

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	go longTask(ctx)
	time.Sleep(4 * time.Second)
}

func longTask(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task finished!")
	case <-ctx.Done():
		fmt.Println("Task cancelled:", ctx.Err())
	}
}
```

## Day 6

### Testing

- **Topics Covered:**
  - Understand testing file of go
  - Your first Go \_test file
  - Table Driven Test
  - Benchmark Test

_Example Code_

```go
func TestAdd_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"2+2", 2, 2, 4},
		{"-1+1", -1, 1, 0},
		{"54+34", 54, 34, 88},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Expected %d got %d", tt.expected, result)
			}
		})
	}
}

```

## Day 7

### Manual testing for CLITodo App

- **Topics Covered:**
  - Write test file manually

_Example Code_

```go
func TestMarkDone(t *testing.T) {
	var t1 Todo
	t1.Add("Walk the dog")
	t1.MarkDone()

	if !t1.status {
		t.Errorf("Expected true got false")
	}
}

```

## Day 8

### Introduction to "net/http"

- **Topics Covered:**
  - Hello Go Web
  - Separate function with handler
  - Multiple routes
  - Query Reader

_Example Code_

```go
func main() {

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)

	fmt.Println("Server running at http://localhost:8834")
	http.ListenAndServe(":8834", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Home!")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is About Page!")
}

```

## Day 9

### Middleware

- **Topics Covered:**
  - Understanding Middleware
  - Logging Middleware
  - Panic Handler Middleware

_Example Code_

```go
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := recover(); err != nil {
			log.Println("Recover from panic...", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			//Server didn't crash, it just logged out panic
		}
		next.ServeHTTP(w, r)
	})
}

```

## Day 10

### "encoding/json"

- **Topics Covered:**
  - Understanding how encoding works
  - Converting JSON to Go struct
  - Handling unknwon JSON
  - Write and read .json file (encode/decode)

_Example Code_

```go
file, _ := os.Create("todo.json")
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	todo := Todo3{ID: 2, Task: "Write and read file"}
	encoder.Encode(todo)

	file2, _ := os.Open("todo.json")
	defer file2.Close()

```

## Day 11

### Routing with Chi/Echo

- **Topics Covered:**
  - Todo API with Chi
  - Todo API with Echo

_Example Code_

```go
r := chi.NewRouter()

r.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
})
//////////////////////////////////////////////////////////////////

e := echo.New()

e.GET("todos", func(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
})

```

## Day 12

### Simple POST/GET Todo API

- **Topics Covered:**
  - GET API
  - POST new todo API
  - GET with id API

_Example Code_

```go
	e.POST("/todos", func(c echo.Context) error {
			var newTodo Todo
			if err := c.Bind(&newTodo); err != nil {
				return c.JSON(http.StatusBadGateway, map[string]string{
					"error": "Not Available",
				})
			}

			newTodo.ID = len(todos) + 1
			todos = append(todos, newTodo)
			return c.JSON(http.StatusCreated, newTodo)
		})

```

## Day 13

### Middleware and Validation with echo

- **Topics Covered:**
  - customErrorHandler with echo

_Example Code_

```go
	e := echo.New()

	e.HTTPErrorHandler = customErrorHander

	e.GET("/todos", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todos)
	})
```
