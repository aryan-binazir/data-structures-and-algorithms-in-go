package main

import "fmt"

func main() {
	fmt.Println(findFirstNonDup("minimum")) // "n"
	fmt.Println(findFirstNonDup("miuimnm")) // "u"
}

func findFirstNonDup(inputStr string) string {
	charMapCount := make(map[string]int)
	returnChar := ""

	for _, char := range inputStr {
		currentChar := string(char)
		if charMapCount[currentChar] != 0 {
			charMapCount[currentChar] += 1
		} else {
			charMapCount[currentChar] = 1
		}
	}

	for _, char := range inputStr {
		currentChar := string(char)

		if charMapCount[currentChar] == 1 {
			returnChar = currentChar
			break
		}
	}

	return returnChar
}
