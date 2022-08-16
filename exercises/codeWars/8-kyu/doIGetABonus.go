package main

import (
	"fmt"
	"strconv"
)

func BonusTime(salary int, bonus bool) string {
	BonusMultiplier := 10
	totalComp := salary

	if bonus == true {
		totalComp = salary * BonusMultiplier
	}

	return ("£" + strconv.Itoa(totalComp))
}

func main() {
	fmt.Println(BonusTime(100, true))  //1000
	fmt.Println(BonusTime(100, false)) //100
}
