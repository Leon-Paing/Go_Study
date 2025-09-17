//This is the example for context used with timeout for presentation.
//Set up 2 seconds time out for context and then we pretend the work takes 3 seconds, so it is closed before finished because context is closed.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	go longTask(ctx)
	time.Sleep(4 * time.Second)
}

func longTask(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task finished!")
	case <-ctx.Done():
		fmt.Println("Task cancelled:", ctx.Err())
	}
}
