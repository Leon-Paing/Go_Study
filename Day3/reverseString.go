package main

import "fmt"

func main() {

	fmt.Println("Reverse string :", Reverse("Hello"))
	fmt.Println("Reverse string :", Reverse("မင်္ဂလာပါ"))
}

func Reverse(s string) string {
	runes := []rune(s)

	//i is 0 initial, j is length-1 , loop will go until i cross j, i is increased by 1 every loop and j is decreased by 1 every loop. Then indexes are reversed and resulting in reversed slice(array).
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
