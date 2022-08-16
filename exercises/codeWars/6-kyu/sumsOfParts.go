package main

import "fmt"

func main() {
	fmt.Println(PartsSums([]uint64{}))                                                                // []uint64{0}
	fmt.Println(PartsSums([]uint64{0, 1, 3, 6, 10}))                                                  // []uint64{20, 20, 19, 16, 10, 0}
	fmt.Println(PartsSums([]uint64{1, 2, 3, 4, 5, 6}))                                                // []uint64{21, 20, 18, 15, 11, 6, 0}
	fmt.Println(PartsSums([]uint64{744125, 935, 407, 454, 430, 90, 144, 6710213, 889, 810, 2579358})) // []uint64{10037855, 9293730, 9292795, 9292388, 9291934, 9291504, 9291414, 9291270, 2581057, 2580168, 2579358, 00}
}

func PartsSums(ls []uint64) []uint64 {
	var totals = []uint64{}

	for index := 0; index <= len(ls); index++ {
		totals = append(totals, sumOfSlice(ls[index:]))
	}
	return totals
}

func sumOfSlice(numbers []uint64) uint64 {
	var total uint64 = 0

	for _, number := range numbers {
		total += number
	}
	return total
}
