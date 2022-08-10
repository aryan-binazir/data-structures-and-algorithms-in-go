package main

import (
	"fmt"
	"strings"
)

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) push(element T) {
	s.data = append(s.data, element)
}

func (s *Stack[T]) read() any {
	if len(s.data) == 0 {
		return nil
	}

	return &s.data[len(s.data)-1]
}

func (s *Stack[T]) pop() *T {
	if len(s.data) == 0 {
		return nil
	}
	index := len(s.data) - 1
	toReturn := s.data[index]
	s.data = s.data[:index]

	return &toReturn
}

func newStack[T any](items T) *Stack[T] {
	return &(Stack[T]{data: []T{items}})
}

func main() {

	fmt.Println("Reversed string is", reverseString("edcba")) // abcde
}

func reverseString(str string) string {
	splitString := strings.Split(str, "")
	s := newStack("")
	for _, value := range splitString {
		s.push(string(value))
	}

	reversedString := ""

	for s.read() != nil {
		currentChar := *s.pop()
		reversedString += string(currentChar)
	}

	return reversedString
}
