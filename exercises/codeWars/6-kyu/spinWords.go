package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(spinWords("Hey fellow warriors"))  //"Hey wollef sroirraw"
	fmt.Println(spinWords("This is a test"))       //"This is a test"
	fmt.Println(spinWords("This is another test")) //This is rehtona test"
}

func spinWords(str string) string {
	spinSlice := []string{}
	splitArr := strings.Split(str, " ")
	for _, word := range splitArr {
		if lengthGreaterThanFour(word) {
			spinSlice = append(spinSlice, reverseWord(word))
		} else {
			spinSlice = append(spinSlice, word)
		}
	}

	return strings.Join(spinSlice, " ")
}

func lengthGreaterThanFour(word string) bool {
	if len(word) > 4 {
		return true
	}
	return false
}

func reverseWord(word string) string {
	returnStr := ""
	wordLength := len(word)

	for index := wordLength - 1; index >= 0; index-- {
		returnStr = returnStr + string(word[index])
	}
	return returnStr
}
