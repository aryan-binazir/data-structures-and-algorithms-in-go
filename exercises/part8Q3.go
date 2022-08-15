package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMissingLetter("the quick brown box jumps over a lazy dog")) // "f"
}

func findMissingLetter(str string) string {
	returnChar := ""

	var alphabetMap = map[string]bool{}

	for _, char := range str {
		currentChar := string(char)
		alphabetMap[currentChar] = true
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for index := 0; index < len(alphabet); index++ {
		if !alphabetMap[string(alphabet[index])] {
			returnChar = string(alphabet[index])
		}
	}
	return returnChar
}
