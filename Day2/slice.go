// A slice is a dynamic array — most commonly used collection in Go.

package main

import "fmt"

func main() {
	nums := []int{}
	nums = append(nums, 3, 4, 5, 6)

	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	for i, v := range nums {
		fmt.Println("index %d = %d/n", i, v)
	}
}
