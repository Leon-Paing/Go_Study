// A struct is like a custom data type that groups related fields together.

package main

import "fmt"

type Car struct {
	Brand  string
	Year   int
	Engine string
}

func main() {
	c1 := Car{Brand: "Mercedes", Year: 2022, Engine: "V8"}
	fmt.Println("This car is a", c1.Year, c1.Brand, ". And this car has", c1.Engine, "engine.")

	c2 := new(Car)
	c2.Brand = "Ferrari"
	c2.Year = 2020
	c2.Engine = "V10"

	fmt.Println(c2)
}
