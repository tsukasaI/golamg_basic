package main

import (
	"fmt"
)

func main() {
	max, min := maxmin(1, 3, 5, 67, 8, 4)
	fmt.Println(max, min)
}

func maxmin(values ...int) (int, int) {
	max := -999999
	min := 999999

	for _, val := range values {
		if max < val {
			max = val
		}
		if min > val {
			min = val
		}
	}
	return max, min
}
