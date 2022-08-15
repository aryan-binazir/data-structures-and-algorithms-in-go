package main

import "fmt"

type key struct {
	rows    int
	columns int
}

func main() {
	fmt.Println(uniquePaths(2, 4, map[key]int{}, map[key]int{}))   // 4
	fmt.Println(uniquePaths(5, 10, map[key]int{}, map[key]int{}))  // 715
	fmt.Println(uniquePaths(10, 20, map[key]int{}, map[key]int{})) // 6906900
	fmt.Println(uniquePaths(30, 50, map[key]int{}, map[key]int{})) // 2627948106712412376                            // Boom
}

func uniquePaths(rows int, columns int, rowsMemo map[key]int, columnsMemo map[key]int) int {
	if rows == 1 || columns == 1 {
		return 1
	}

	_, ok := rowsMemo[key{rows, columns}]

	if !ok {
		rowsMemo[key{rows, columns}] = uniquePaths(rows-1, columns, rowsMemo, columnsMemo) +
			uniquePaths(rows, columns-1, rowsMemo, columnsMemo)
	}

	return rowsMemo[key{rows, columns}]
}
