package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(removeIntervals([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
	// [[1,6],[8,10],[15,18]]
	fmt.Println(removeIntervals([][]int{[]int{1, 4}, []int{4, 5}}))
	//[[1,5]]
	fmt.Println(removeIntervals([][]int{[]int{4, 5}}))
	//[[4,5]]
	fmt.Println(removeIntervals([][]int{[]int{1, 3}, []int{2, 7}, []int{3, 5}, []int{4, 6}, []int{14, 21}}))
	//[1,7] [14,21]
}

func removeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	intervalsDup := make([][]int, len(intervals))
	copy(intervalsDup, intervals)
	sort.Slice(intervalsDup, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	currentIndex := 0

	for currentIndex <= len(intervalsDup) {
		currentSl := intervalsDup[0]
		nextSl := intervalsDup[1]
		if currentSl[1] > nextSl[0] {
			intervalsDup = intervalsDup[2:]
			newIntervals := [][]int{[]int{currentSl[0], nextSl[1]}}
			for _, slice := range intervalsDup {
				newIntervals = append(newIntervals, slice)
			}
			intervalsDup = newIntervals
		} else {
			currentIndex++
		}

	}
	return intervalsDup
}
