package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solution("abc"))    // ['ab', 'c_']
	fmt.Println(solution("abcdef")) //['ab', 'cd', 'ef']
}

func solution(str string) []string {
	strArr := strings.Split(str, "")
	resultArr := []string{}

	for len(strArr) > 0 {
		if len(strArr) >= 2 {
			resultArr = append(resultArr, getPair(strArr[:2]))
			strArr = strArr[2:]
		} else {
			resultArr = append(resultArr, strArr[0]+"_")
			strArr = []string{}
		}
	}
	return resultArr
}

func getPair(chars []string) string {
	return strings.Join(chars, "")
}
