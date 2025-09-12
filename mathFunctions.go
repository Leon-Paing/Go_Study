package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	fmt.Println("Sum value:", add(2, 4))

	x, y := divide(4, 6)
	fmt.Println("Quotient:", x, ",Remainder:", y)
}
