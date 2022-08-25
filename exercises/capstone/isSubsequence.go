package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) //true
	fmt.Println(isSubsequence("axc", "ahbgdc")) //false
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return false
	}

	i, j := 0, 0
	var result bool
	lenS, lenT := len(s), len(t)

	for true {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}

		if i > lenS-1 {
			result = true
			break
		}
		if j > lenT-1 {
			result = false
			break
		}
	}

	return result
}
