// errors.Is → check if an error is a specific error.
// errors.As → check if error is of a specific type.

package main

import (
	"errors"
	"fmt"
)

func main() {
	err := findUser(1)

	if errors.Is(err, ErrorNotFound) {
		fmt.Println("User not found in database.")
	} else {
		fmt.Println("Other error occured!")
	}
}

var ErrorNotFound = errors.New("not found.")

func findUser(id int) error {
	return fmt.Errorf("database error %w", ErrorNotFound)
}
