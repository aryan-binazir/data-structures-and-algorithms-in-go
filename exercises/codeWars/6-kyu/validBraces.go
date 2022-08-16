package main

import "fmt"

func ValidBraces(str string) bool {

	openers := []string{"(", "{", "["}
	closers := []string{")", "}", "]"}
	closersMap := make(map[string]string)
	closersMap["]"] = "["
	closersMap[")"] = "("
	closersMap["}"] = "{"
	currentQueue := []string{}

	for _, value := range str {
		currentBracket := string(value)

		if contains(openers, currentBracket) {
			currentQueue = append(currentQueue, currentBracket)
		} else {
			if !contains(closers, currentBracket) {
				continue
			} else {
				if currentQueue[len(currentQueue)-1] != closersMap[currentBracket] {
					return false
				} else {
					currentQueue = currentQueue[:len(currentQueue)-1]
				}
			}
		}
	}

	if len(currentQueue) == 0 {
		return true
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

func main() {
	fmt.Println(ValidBraces("(){}[]"))  //true
	fmt.Println(ValidBraces("([{}])"))  //true
	fmt.Println(ValidBraces("(}"))      //false
	fmt.Println(ValidBraces("[(])"))    //false
	fmt.Println(ValidBraces("[({)](]")) //false
}
