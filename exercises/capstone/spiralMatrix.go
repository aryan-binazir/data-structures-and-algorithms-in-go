/*
Problem:

Input: m x n matrix
Output: all elements in spiral order

Rules:
- m == matrix.length
- n == matrix[i].length
- 1 <= m, n <= 10
- 100 <= matrix[i][j] <= 100

Examples:
Example 1:

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:

Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

Example 3:

Data Structures:
-Nested arrays
Algorithm:
back to front if smaller then subtract from total

Algorithm
n is length of outer array (3)
m is length of inner array (3)

Loop
	- from [0,0] to [m-1, 0] add to result
	- from [m-1, 0] to [m-1, n-1] add to result
	- from [m-1, n-1] to [0, n-1]
	- from [0, n-1] to [0, ]

- add first nested array values to `resultArr`
- add [m-1, 0] and [m-1,1] up to [m-1, n-1] to `resultArr` where n is number of arrays nested
- reverse above from [m-1, n-1] to [0, n-1]
- repeat above from [0, n-1] to [0,]
*/

// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]

package main

import "fmt"

// arr1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// Incomplete

func main() {
	fmt.Println(spiral([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	// [1,2,3,4,8,12,11,10,9,5,6,7]
	// fmt.Println(spiral([]interface{}{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
	// [1,2,3,6,9,8,7,4,5]
}

func spiral(matrix [][]int) []int {
	// initialize result
	// loop until array empty
	// remove first array and add to result
	// loop through rest of arrays and add last element to result
	// remove last array and add in reverse order to result
	// add first element of first array to result

	result := []int{}
	for len(matrix) > 0 {
		for _, val := range matrix[0] {
			result = append(result, val)
		}
		matrix = matrix[1:]
		// 	matrix = matrix[1:]
		// 	if len(matrix) == 0 {
		// 		break
		// 	}
		// 	for _, slice := range matrix {
		// 		result = append(result, matrix[len(slice)-1])
		// 		// matrix = matrix[0 : len(slice)-1]
		// 	}
		// 	if len(matrix) == 0 {
		// 		break
		// 	}

	}
	return result
}
