// A map is a key–value store (like Python dict / JS object).
//When looped through a map, the order is random

package main

import "fmt"

func main() {
	scores := make(map[string]int)

	scores["YanPai"] = 100
	scores["Zuck"] = 0

	fmt.Println("Score of Yan Pai is:", scores["YanPai"])

	val, ok := scores["Musk"]

	if ok {
		fmt.Println("Score of Charlie is :", val)
	} else {
		fmt.Println("Charlie is not found!")
	}

	for name, score := range scores {
		fmt.Println("Score of", name, "is", score)
	}
}
