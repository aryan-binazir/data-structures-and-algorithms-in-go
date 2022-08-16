package main

import "fmt"

func main() {
	slice := []interface{}{
		1,
		2,
		3,
		[]interface{}{4, 5, 6},
		7,
		[]interface{}{8,
			[]interface{}{9, 10, 11,
				[]interface{}{12, 13, 14},
			},
		},
		[]interface{}{15, 16, 17, 18, 19,
			[]interface{}{20, 21, 22,
				[]interface{}{23, 24, 25,
					[]interface{}{26, 27, 29},
				}, 30, 31,
			}, 32,
		}, 33,
	}
	printNum(slice)
}

func printNum(slice []interface{}) {
	if len(slice) == 0 {
		return
	}

	value, ok := slice[0].(int)

	if ok {
		fmt.Println(value)
		printNum(slice[1:])
	} else {
		value, ok := slice[0].([]interface{})

		if ok {
			printNum(value)
			printNum(slice[1:])
		}
	}
}
