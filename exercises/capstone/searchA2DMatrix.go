// complete
package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 60}}, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	found := false
	left, right := 0, len(matrix)-1
	for left <= right {
		mid := left + ((right - left) / 2)
		fmt.Println(mid)
		if (matrix[mid][len(matrix[mid])-1] >= target) && matrix[mid][0] <= target {
			searchSubarray(matrix[mid], target)
		} else if target < matrix[mid][len(matrix[mid])-1] {
			right = mid + 1
		} else {
			left = mid - 1
		}
	}

	return found
}

func searchSubarray(subarray []int, target int) bool {
	found := false

	left, right := 0, len(subarray)-1
	for left <= right {
		mid := left + ((right - left) / 2)
		if subarray[mid] == target {
			found = true
			break
		} else if target > subarray[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return found
}

/*

 */
