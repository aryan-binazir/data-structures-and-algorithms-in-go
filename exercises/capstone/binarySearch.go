package main

import "fmt"

func main() {
	fmt.Println(findNumber([]int{1, 5, 8, 9, 22, 55, 345, 3466, 24545, 45362526, 5625626246}, 345))                    // true
	fmt.Println(findNumber([]int{1, 5, 8, 9, 22, 55, 345, 3466, 24545, 45362526, 5625626246}, 3451212))                // false
	fmt.Println(findNumber([]int{-44, -12, -4, -2, 0, 1, 5, 8, 9, 22, 55, 345, 3466, 24545, 45362526, 5625626246}, 0)) // true
	fmt.Println(findNumber([]int{32345}, 32345))                                                                       // true
	fmt.Println(findNumber([]int{123}, 32345))                                                                         // false
	fmt.Println(findNumber([]int{}, 32345))                                                                            // false
}

func findNumber(slice []int, target int) bool {

	if len(slice) == 0 {
		return false
	} else if slice[0] == target {
		return true
	} else if len(slice) == 1 {
		return false
	}

	middleIndex := len(slice) / 2
	middleValue := slice[middleIndex]
	var found bool

	if middleValue == target {
		return true
	} else if middleValue > target {
		found = findNumber(slice[:middleIndex], target)
	} else {
		found = findNumber(slice[middleIndex:], target)
	}
	return found
}
