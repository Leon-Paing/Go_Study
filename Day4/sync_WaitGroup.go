// Wait for goroutines to finish without blocking with Sleep.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done! Exiting....")
	time.Sleep(1 * time.Second)
	fmt.Println("\nClosed Program!")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker", id, "starting...")
	time.Sleep(1 * time.Second)
	fmt.Println("Worker", id, "Done!")
}
