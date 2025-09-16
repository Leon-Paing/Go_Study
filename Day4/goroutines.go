// In Go, every program runs in a main goroutine by default.
// A goroutine is a lightweight thread of execution managed by the Go runtime.
// Start a goroutine with the go keyword.

package main

import (
	"fmt"
	"time"
)

// go printNumbers() launches the function concurrently.

// If main ends, all goroutines stop. That’s why we added time.Sleep.

// Output interleaves because goroutines run in parallel.
func main() {

	go printNumbers()
	go printHello()

	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
}

func printNumbers() {
	for i := 0; i <= 6; i++ {
		fmt.Println("Number:", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func printHello() {
	for i := 0; i <= 6; i++ {
		fmt.Println("Hello!", i)
		time.Sleep(300 * time.Millisecond)
	}
}
