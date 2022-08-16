package main

import "fmt"

func main() {
	fmt.Println(findOdd([]int{7}))                                     //7
	fmt.Println(findOdd([]int{0}))                                     //0
	fmt.Println(findOdd([]int{1, 1, 2}))                               //2
	fmt.Println(findOdd([]int{0, 1, 0, 1, 0}))                         //0
	fmt.Println(findOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})) //4
}

func findOdd(seq []int) int {
	oddMap := make(map[int]bool)

	for _, number := range seq {
		if oddMap[number] {
			oddMap[number] = !oddMap[number]
		} else {
			oddMap[number] = true
		}
	}

	var returnVal int

	for key, value := range oddMap {
		if value == true {
			returnVal = key
		}
	}
	return returnVal
}
