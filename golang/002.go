package main

import (
	"fmt"
	"math"
)

func findMaxMin(arr []int) []int {
	var max, min = 0, math.MaxInt64
	for _, ch := range arr {
		if ch > max {
			max = ch
		}
		if ch < min {
			min = ch
		}
	}
	return []int{max, min}
}

func main() {
	fmt.Println(findMaxMin([]int{1, 4, 7, 2, 3, 4, 3, 6, 8}))
}
