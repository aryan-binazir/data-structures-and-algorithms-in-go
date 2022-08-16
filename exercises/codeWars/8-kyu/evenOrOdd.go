package main

import "fmt"

func evenOrOdd(number int) string {
	var returnVal string

	if number%2 == 0 {
		returnVal = "Even"
	} else {
		returnVal = "Odd"
	}
	return returnVal
}

func main() {
	fmt.Println(evenOrOdd(2)) // Even
	fmt.Println(evenOrOdd(3)) // Odd
}
