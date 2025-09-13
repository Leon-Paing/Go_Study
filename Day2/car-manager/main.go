package main

import "fmt"

func main() {

	fmt.Println("-------------\nCar Showroom\n-------------")

	car1 := Car{Brand: "Toyota", Year: 2024, Price: 300.453}
	fmt.Println(car1.Detail())
	fmt.Println("Discounted Price:", car1.Discount(10), "$\n")

	//create cars slice
	cars := []Car{
		{Brand: "Toyota", Year: 2020, Price: 290.87},
		{Brand: "Honda", Year: 2021, Price: 280.34},
		{Brand: "Subaru", Year: 2022, Price: 340.97},
	}

	for _, c := range cars {
		fmt.Println(c.Detail())
		fmt.Println("Discount Price for", c.Brand, c.Discount(2), "\n")
	}

	//create cars slice to carsMap map key value pair (key is Brand)
	carsMap := make(map[string]Car)
	for _, c := range cars {
		carsMap[c.Brand] = c
	}

	mitsubishi, ok := carsMap["Mitsubishi"]
	if ok {
		fmt.Println("Car brand found: ", mitsubishi.Brand)
	} else {
		fmt.Println("Car brand NOT FOUND!\n")
	}

	//interface
	var fourweel Valuable

	fourweel = Car{Brand: "Ford", Year: 2015, Price: 380.234}
	fmt.Println("Value is", fourweel.Value())
}

type Car struct {
	Brand string
	Year  int
	Price float64
}

func (c Car) Detail() (string, int, float64) {
	return c.Brand, c.Year, c.Price
}

func (c Car) Discount(percent float64) float64 {
	return c.Price - ((percent * 100) / c.Price)
}

func (c Car) Value() float64 {
	return c.Price
}

type Valuable interface {
	Value() float64
}
