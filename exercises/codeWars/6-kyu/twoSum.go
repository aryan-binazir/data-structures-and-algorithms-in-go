package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 3}, 4))              // [2]int{0, 2}
	fmt.Println(twoSum([]int{1234, 5678, 9012}, 14690)) // [2]int{1, 2}
	fmt.Println(twoSum([]int{2, 2, 3}, 4))              // [2]int{0, 1}
}

func twoSum(numbers []int, target int) [2]int {
	var addToTuple = []int{}

	for index1 := 0; index1 < len(numbers)-1; index1++ {
		for index2 := index1 + 1; index2 < len(numbers); index2++ {
			if (numbers[index1] + numbers[index2]) == target {
				addToTuple = append(addToTuple, index1)
				addToTuple = append(addToTuple, index2)
				break
			}
		}
		if len(addToTuple) != 0 {
			break
		}
	}

	return [2]int{addToTuple[0], addToTuple[1]}
}
