// You can create formatted error messages using fmt.Errorf:
package main

import "fmt"

func main() {
	err := readFile("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func readFile(file string) error {
	return fmt.Errorf("error reading file %s", file)
}
