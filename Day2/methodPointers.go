// Methods are just functions with a receiver.
// They let you attach behavior to structs (similar to OOP).

package main

import "fmt"

func main() {
	rect := Rectangle{Width: 3, Height: 4}
	fmt.Println("Before modified rectangle area is", rect.Area())

	rect.Scale(3)
	fmt.Println("After modified rectangle area is", rect.Area())
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}

// Use pointer receivers when the method modifies the struct.
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}
