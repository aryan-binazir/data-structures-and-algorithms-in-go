func minimumTotal(triangle [][]int) int {
	cache := map[[2]int]int{}
	return findMinPathSum(triangle, 0, 0, cache) // pass row, column and cache
}

func findMinPathSum(triangle [][]int, row, column int, cache map[[2]int]int) int {
	if row == len(triangle)-1 { // Base case when at last row
		return triangle[row][column]
	}

	if val, ok := cache[[2]int{row, column}]; ok { // if exists in cache, return value
		return val
	}

	currValue := triangle[row][column]
	firstSum := findMinPathSum(triangle, row+1, column, cache)    //recurse with same column
	secondSum := findMinPathSum(triangle, row+1, column+1, cache) // recurse with column + 1

	var sumToAdd int

	if firstSum > secondSum { // add current value to minimum sum of returned recursive functions and return
		sumToAdd = currValue + secondSum
	} else {
		sumToAdd = currValue + firstSum
	}

	cache[[2]int{row, column}] = sumToAdd
	return cache[[2]int{row, column}]
}


