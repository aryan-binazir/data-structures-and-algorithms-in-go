package main

import "fmt"

func main() {
	fmt.Println(arrayIntersection([]int{1, 2, 3, 4, 5}, []int{0, 2, 4, 6, 8})) // [2, 4]
}

func arrayIntersection(arr1 []int, arr2 []int) []int {
	arr1Map := make(map[int]bool)
	result := []int{}

	for _, num := range arr1 {
		arr1Map[num] = true
	}

	for _, num := range arr2 {
		if arr1Map[num] {
			result = append(result, num)
		}
	}

	return result
}
