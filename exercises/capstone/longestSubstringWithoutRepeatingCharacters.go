package main // almost done some bugs

import "fmt"

func main() {
	// fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	// fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew")) // 3
}

func lengthOfLongestSubstring(s string) int {

	// - map charMap initialize
	// - set anchor and runner to 0
	// - set maxSubstring to 0
	// - add first char to charMap

	// - loop until runner out of bounds
	// 	- move runner by one
	// 		- if repeated then move anchor to previous spot and set new position in hash
	// 	- if not repeated check if size (runner - anchor + 1) is greater than maxSubstring
	// 		- if it is reassign maxSubstring to new length
	// - return maxSubstring

	charMap := make(map[byte]int)
	anchor, runner, maxSubstring := 0, 1, 1
	charMap[s[anchor]] = anchor

	for runner < len(s) {

		// fmt.Println(string(s[runner]))
		if _, ok := charMap[s[runner]]; ok {
			// fmt.Println("here", charMap[(s[runner]])
			// fmt.Println("charMap", charMap)
			anchor = charMap[s[runner]] + 1
			charMap[s[anchor]] = anchor
		} else {
			newLength := runner - anchor + 1
			if newLength > maxSubstring {
				maxSubstring = newLength
			}
		}
		runner++
		fmt.Println(anchor, runner)

	}

	return maxSubstring
}

/*
anchor and runner start at 0
anchor moves when repeat char
runner moves when no repeat

runner needs to keep list of keys of letters and index, also check each move

ends when runner goes out of bounds

- map charMap initialize
- set anchor and runner to 0
- set maxSubstring to 0
- add first char to charMap
- loop until runner out of bounds
	- move runner by one
		- if repeated then move anchor to previous spot + 1 and set new position in hash
	- if not repeated check if size (runner - anchor + 1) is greater than maxSubstring
		- if it is reassign maxSubstring to new length
- return maxSubstring

*/
