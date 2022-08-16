package main

import "fmt"

func solution(word string) string {
	runes := []rune(word)
	result := ""

	for index := len(word) - 1; index >= 0; index-- {
		result = result + string(runes[index])
	}
	return result
}

func main() {
	fmt.Println(solution("world")) // 'dlrow'
	fmt.Println(solution("word"))  //'drow'
}
