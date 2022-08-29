package main

import "fmt"

func minPathSum(grid [][]int) int {
	cache := map[[2]int]int{}
	row, column := 0, 0
	maxRow := len(grid) - 1
	maxColumn := len(grid[0]) - 1
	return helper(grid, row, column, maxRow, maxColumn, cache)
}

func helper(grid [][]int, row, col, maxRow, maxColumn int, cache map[[2]int]int) int {
	if val, ok := cache[[2]int{row, col}]; ok {
		return val
	}
	fmt.Println(row, col)

	result := []int{}
	if row == maxRow && col == maxColumn {
		return grid[row][col]
	}
	if row < maxRow { // 1
		result = append(result, grid[row][col]+helper(grid, row+1, col, maxRow, maxColumn, cache))
	}
	if col < maxColumn { // 2
		result = append(result, grid[row][col]+helper(grid, row, col+1, maxRow, maxColumn, cache))
	}
	if len(result) == 2 {
		if result[0] > result[1] {
			cache[[2]int{row, col}] = result[1]
			return cache[[2]int{row, col}]
		}
		cache[[2]int{row, col}] = result[0]
		return result[0]
	} else {
		cache[[2]int{row, col}] = result[0]
		return result[0]
	}
}

func main() {

	// arr1 := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	// fmt.Println(minPathSum(arr1))

	arr2 := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(minPathSum(arr2))
}
