package main

import (
	"errors"
	"fmt"
)

func main() {

	st := Stack{}

	st.Push(10)
	st.Push(20)
	st.Push(30)
	fmt.Println(st)

	remVal, err := st.Pop()
	if err != nil {
		fmt.Println("Error%w:", err)
	} else {
		fmt.Println("Removed:", remVal)
	}

	resl, err := st.Peek()
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Top element is:", resl)
	}

	fmt.Println(st)

}

type Stack struct {
	stack []int
}

func (s *Stack) Push(x int) {
	s.stack = append(s.stack, x)
}

func (s *Stack) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("Stack is empty")
	}
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return val, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("Stack is empty")
	}
	return s.stack[len(s.stack)-1], nil
}
