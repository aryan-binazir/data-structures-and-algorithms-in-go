package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(order("is2 Thi1s T4est 3a") == "Thi1s is2 3a T4est")
	fmt.Println(order("4of Fo1r pe6ople g3ood th5e the2") == "Fo1r the2 g3ood 4of th5e pe6ople")
	fmt.Println(order("") == "")
}

func order(sentence string) string {
	if sentence == "" {
		return ""
	}
	result := []string{}

	wordMap := make(map[int]string)
	sentenceArr := strings.Split(sentence, " ")

	for _, word := range sentenceArr {
		wordMap[findNum(word)] = word
	}

	for index := 1; index <= len(sentenceArr); index++ {
		result = append(result, wordMap[index])
	}

	return strings.Join(result, " ")
}

func findNum(word string) int {
	numArr := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var returnNum int

	for _, value := range word {
		if contains(numArr, string(value)) {
			returnNum, _ = strconv.Atoi(string(value))
			break
		}
	}
	return returnNum
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
