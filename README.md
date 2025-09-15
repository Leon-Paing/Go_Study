# 🦁 Learn Go Programming with Me

### This is the Go Programming guide for **beginners to experts**.  

> 💡 Remember: *Consistency is Key!*

<p align="center">
	<img src="https://media2.giphy.com/media/v1.Y2lkPTc5MGI3NjExaHo3OTg3b3Qya2F5NGFjY3lmODcya3lncDRraW1qOTZmbmRmMmMzMyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/PhTSmzCqkliqIJ9ZtZ/giphy.gif">
</p>

---

## 📚 Table of Contents

[Day 1](#day-1) | Introduction to Go & Setup <br>
[Day 2](#day-2) | Dealing with data <br>
[Day 3](#day-3) | custom Algorithms <br>
---

## Day 1 
### Introduction

- **Topics Covered:**  
  - Installing Go.  
  - Learn go run, go build, go mod init.
  - Review variables, types, loops, functions. 
  - Writing your first Go program.

*Example Code*
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
  - Study structs, slices, maps, methods, interfaces.
  - Mini project named car-manager.

*Example Code*
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
  - Implement custom made stack and reverse string.
  - Mini project named TodoCLI.

*Example Code*
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
