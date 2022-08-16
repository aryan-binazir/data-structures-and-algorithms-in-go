package main

import "fmt"

func removeChar(word string) string {
	runes := []rune(word)
	result := ""

	for index := 1; index < len(word)-1; index++ {
		result = result + string(runes[index])
	}
	return result
}

func main() {
	fmt.Println(removeChar("country"))   //ountr
	fmt.Println(removeChar("dancehall")) //ancehal
}
