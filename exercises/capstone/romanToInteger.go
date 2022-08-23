package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("III"))
	//Output: 3
	fmt.Println(romanToInt("LVIII"))
	//Output: 58
	fmt.Println(romanToInt("MCMXCIV"))
	//Output: 1994
}

func romanToInt(s string) int {
	var values = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	numArr := strings.Split(s, "")
	result := values[numArr[len(numArr)-1]]

	for i := len(numArr) - 2; i >= 0; i-- {
		currentNumeral, previousNumeral := numArr[i], numArr[i+1]

		if values[currentNumeral] < values[previousNumeral] {
			result -= values[currentNumeral]
		} else {
			result += values[currentNumeral]
		}

	}
	return result
}
