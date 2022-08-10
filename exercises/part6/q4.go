package main

import "fmt"

func main() {
	fmt.Println(containsX("hasXdontyousee"))   // true
	fmt.Println(containsX("hasnoxdontyousee")) // false
}

// Improve efficiency for best- & average-case scenarios
func containsX(str string) bool {

	for index := 0; index < len(str); index++ {
		if string(str[index]) == "X" {
			return true
		}
	}
	return false
}
