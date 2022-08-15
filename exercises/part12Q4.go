package main

import "fmt"

func main() {
	fmt.Println(findXIndex("abcdefghijklmnopqrstuvwxyz")) // 23
	fmt.Println(findXIndex("ajsufxis"))                   // 5
}

func findXIndex(str string) int {
	if string(str[0]) == "x" {
		return 0
	}

	return 1 + findXIndex(str[1:])
}
