package main

import (
	"fmt"
	"strings"
)

var alphabetScore = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
}

func main() {
	fmt.Println(high("man i need a taxi up to ubud"))             //"taxi"
	fmt.Println(high("what time are we climbing up the volcano")) //"volcano"
}

func high(s string) string {
	wordArr := strings.Split(s, " ")
	highestWord := ""
	highestWordScore := 0

	for _, word := range wordArr {
		wordScore := findScore(word)
		if wordScore > highestWordScore {
			highestWord = word
			highestWordScore = wordScore
		}
	}
	return highestWord
}

func findScore(word string) int {
	totalScore := 0

	for _, char := range word {
		totalScore += alphabetScore[string(char)]
	}
	return totalScore
}
