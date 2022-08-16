package main

import "fmt"

func PositiveSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if isPositive(num) {
			sum += num
		}
	}
	return sum
}

func isPositive(num int) bool {
	if num > 0 {
		return true
	}
	return false
}

func main() {
	// Calculate sum of positive numbers in an array
	fmt.Println(PositiveSum([]int{1, 2, 3, 4, 5}))      //15
	fmt.Println(PositiveSum([]int{1, -2, 3, 4, 5}))     //13
	fmt.Println(PositiveSum([]int{}))                   //0
	fmt.Println(PositiveSum([]int{-1, -2, -3, -4, -5})) //0
	fmt.Println(PositiveSum([]int{-1, 2, 3, 4, -5}))    //9
}
