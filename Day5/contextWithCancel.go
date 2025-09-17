//You manually cancel when you want.
//We created function to loop infinitely but with context control. So, function will be looping infinitely until context is cancelled manually.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 2; i++ {
		go worker(ctx, i)
	}
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func worker(ctx context.Context, id int) {
	//infinite loop until ctx is cancelled and break the loop.
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker", id, "stopping...", ctx.Err())
			return
		default:
			fmt.Println("Worker", id, "working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
