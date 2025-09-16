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
[Day 3](#day-3) | Custom algorithms <br>
[Day 4](#day-4) | Go routines and Channels <br>
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
## Day 4
### Go routines and Channels

- **Topics Covered:**  
  - Understand about go routines.
  - Learn why channels are important.
  - Know difference between buffered and unbuffered channels.
  - Channel sequence.
  - Sync and WaitGroup.

*Example Code*
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
