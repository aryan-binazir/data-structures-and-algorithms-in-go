package main

import "fmt"

func main() {
	fmt.Println(reverseString([]string{"h", "e", "l", "l", "o"}))      //["o","l","l","e","h"]
	fmt.Println(reverseString([]string{"H", "a", "n", "n", "a", "h"})) //"h","a","n","n","a","H"

}

func reverseString(s []string) []string {
	start := 0
	end := len(s) - 1

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}

	return s
}
