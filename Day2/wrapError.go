// Go lets you wrap errors so you don’t lose the original cause.

package main

import (
	"errors"
	"fmt"
)

func main() {
	err := initApp()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func readConfig() error {
	return errors.New("file not found!")
}

func initApp() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("init app failed : %w", err)
	}
	return nil
}
