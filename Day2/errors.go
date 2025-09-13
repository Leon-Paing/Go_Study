// In Go, functions often return two values:
// 1.The result
// 2.An error (if something went wrong)

package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := add(3, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

func add(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0!")
	}
	return a + b, nil
}
