package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println("Value of i is:", i)
	}

	n := 1
	for n < 5 {
		fmt.Println("Value of n:", n)
		n++
	}

	for {
		fmt.Println("looping.....")
		break
	}
}
