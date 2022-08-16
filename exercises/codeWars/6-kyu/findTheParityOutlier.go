package main

import "fmt"

func main() {
	fmt.Println(findOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36})) // 11 (the only odd number)
	fmt.Println(findOutlier([]int{160, 3, 1719, 19, 11, 13, -21})) // 160 (the only even number)
}

func findOutlier(integers []int) int {
	evenArr := []int{}
	oddArr := []int{}

	for _, num := range integers {
		if isEven(num) {
			evenArr = append(evenArr, num)
		} else {
			oddArr = append(oddArr, num)
		}

		if (len(evenArr) > 1 && len(oddArr) == 1) ||
			(len(oddArr) > 1 && len(evenArr) == 1) {
			break
		}
	}
	var returnVal int
	if len(evenArr) == 1 {
		returnVal = evenArr[0]
	} else {
		returnVal = oddArr[0]
	}
	return returnVal

	return 5
}

func isEven(number int) bool {
	return number%2 == 0
}
