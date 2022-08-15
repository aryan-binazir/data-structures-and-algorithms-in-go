package main

import "fmt"

func main() {
	fmt.Println(countChars([]string{"12"}))                     // 2
	fmt.Println(countChars([]string{"ab", "c", "def", "ghij"})) // 10
}

func countChars(arrOfChars []string) int {
	if len(arrOfChars) == 1 {
		return len(arrOfChars[0])
	}
	return len(arrOfChars[0]) + countChars(arrOfChars[1:])
}
