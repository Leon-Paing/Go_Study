//context.WithValue() attach key value pairs for data passing.
//It is a common usage in passing metadata like user id, user name, etc.., So you can receive value from context easily.
//!!!Important Notice!!!
//Go documents discourage using context.WithValue() for passing optional parameters like config settings, etc.,
//Use it only for request-scoped data.

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "userID", 23)
	passData(ctx)

}

func passData(ctx context.Context) {
	userID := ctx.Value("userID")
	fmt.Println("Processing request for user", userID)
}
