package main

import (
	"fmt"
	"sort"
)

func main() {
	a1 := []string{"arp", "live", "strong"}
	a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	fmt.Println(inArray(a1, a2)) // ["arp", "live", "strong"]

	b1 := []string{"tarp", "mice", "bull"}
	b2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	fmt.Println(inArray(b1, b2)) // []
}

func inArray(array1 []string, array2 []string) []string {
	result := []string{}
	for _, word1 := range array1 {
		for _, word2 := range array2 {
			if stringInString(word1, word2) && !contains(result, word1) {
				result = append(result, word1)
			}
		}
	}
	sort.Strings(result)

	return result
}

func stringInString(str1 string, str2 string) bool {
	for index := 0; index <= len(str2)-1; index++ {
		for index2 := index; index2 <= len(str2); index2++ {
			if str2[index:index2] == str1 {
				return true
			}
		}
	}
	return false
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
