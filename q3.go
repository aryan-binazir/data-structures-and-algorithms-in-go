// Not complete

package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMissingLetter("the quick brown box jumps over a lazy dog")) // "f"
}

func findMissingLetter(str string) string {
	returnChar := ""

	var alphabetMap = map[string]bool{
		"a": false,
		"b": false,
		"c": false,
		"d": false,
		"e": false,
		"f": false,
		"g": false,
		"h": false,
		"i": false,
		"j": false,
		"k": false,
		"l": false,
		"m": false,
		"n": false,
		"o": false,
		"p": false,
		"q": false,
		"r": false,
		"s": false,
		"t": false,
		"u": false,
		"v": false,
		"w": false,
		"x": false,
		"y": false,
		"z": false,
	}

	keys := []string{}

	for _, char := range str {
		currentChar := string(char)

		if currentChar != " " {
			keys = append(keys, currentChar)
		}

		alphabetMap[currentChar] = true
	}

	for _, key := range keys {
		fmt.Println(key, alphabetMap[key])
		if alphabetMap[key] == false {
			returnChar = key
		}
	}

	return returnChar
}
