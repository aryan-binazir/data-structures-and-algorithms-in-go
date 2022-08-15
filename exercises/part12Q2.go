package main

import "fmt"

func main() {
	fmt.Println(golomb(1, map[int]int{}))    // 1
	fmt.Println(golomb(3, map[int]int{}))    // 2
	fmt.Println(golomb(13, map[int]int{}))   // 6
	fmt.Println(golomb(70, map[int]int{}))   // 17
	fmt.Println(golomb(10, map[int]int{}))   // 5
	fmt.Println(golomb(1000, map[int]int{})) // 86
}

func golomb(num int, memo map[int]int) int {
	if num == 1 {
		return 1
	}

	_, ok := memo[num]
	if !ok {
		memo[num] = 1 + golomb(num-golomb(golomb(num-1, memo), memo), memo)
	}

	return memo[num]
}
