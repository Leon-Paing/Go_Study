// Interfaces define behavior, not data.
// Any type that implements the required methods automatically satisfies the interface (no explicit implements).

package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// I declared Rectangle2 because Rectangle is already declared under same folder(module) and same package

type Rectangle2 struct {
	Width, Height float64
}

func (r Rectangle2) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var s Shape

	s = Circle{Radius: 12}
	fmt.Println("Area of Circle is", s.Area())

	s = Rectangle2{Width: 4, Height: 3}
	fmt.Println("Area of rectangle is", s.Area())
}
