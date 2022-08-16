package main

import "fmt"

func main() {
	fmt.Println(multiple3And5(10)) //23
}

func multiple3And5(number int) int {
	sum := 0
	for count := 1; count < number; count += 1 {
		if multipleOf3(count) || multipleOf5(count) {
			sum += count
		}
	}
	return sum
}

func multipleOf3(number int) bool {
	return number%3 == 0
}

func multipleOf5(number int) bool {
	return number%5 == 0
}
