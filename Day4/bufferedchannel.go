// Sender can send up to capacity values before blocking.
// Buffered channels decouple sender and receiver.
package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20
	// ch <- 30 This will blocked(deadlock) and cause error

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch <- 30 //We can send again because channel is empty now.
	fmt.Println(<-ch)
}
