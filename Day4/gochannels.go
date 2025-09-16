// Goroutines are great, but they need a way to communicate safely.
// That’s what channels are for.

package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	ch := make(chan int)

	go squares(numbers, ch)
	for result := range ch {
		fmt.Println("Result:", result)
	}
}

func squares(numbers []int, ch chan int) {
	for _, n := range numbers {
		ch <- n * n
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
}
