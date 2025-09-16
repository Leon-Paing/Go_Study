// Sender and receiver must “meet” at the same time.

package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello message!"

	}()

	msg := <-ch
	fmt.Println("Received message:", msg)
}
