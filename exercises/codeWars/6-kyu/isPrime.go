package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPrime(1))  // false
	fmt.Println(isPrime(2))  // true
	fmt.Println(isPrime(-1)) // false
	fmt.Println(isPrime(73)) // true
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}

	var sqRoot int = int(math.Sqrt(float64(number)))

	for index := 2; index <= sqRoot; index++ {
		if number%index == 0 {
			return false
		}
	}
	return true
}
