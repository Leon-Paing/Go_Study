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
