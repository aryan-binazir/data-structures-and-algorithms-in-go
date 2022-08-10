package main

import (
	"fmt"
)

func main() {
	s := stack.newStack[int](1)
	fmt.Println(s)
	s.push(6)
	fmt.Println(s) // 1,6
	s.pop()
	fmt.Println(s) // 1
	s.read()
	fmt.Println(s) // 1
}
