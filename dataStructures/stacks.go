package main

import "fmt"

type Stack struct {
	items []int
}

// Push adds value to end of stack
func (s *Stack) Push(toAdd int) {
	s.items = append(s.items, toAdd)
}

// Pop removes value at end of stack and returns value
func (s *Stack) Pop() int {
	itemsLength := len(s.items)
	toRemove := s.items[itemsLength-1]
	s.items = s.items[:itemsLength-1]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
	myStack.Push(400)
	fmt.Println(myStack)
}
