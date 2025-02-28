package main

import (
	"fmt"
)

func zeroSumSubArr(arr []int) [][]int {
	sum := 0

	subs := [][]int{}

	start := 0
	stop := 0

	for i := 1; i < len(arr); i++{
		for j := i; j < len(arr);j++{
			
			sum += arr[j]
			fmt.Println(sum)
			if sum == 0{
				stop = i
				subs = append(subs, []int{start, stop})
				start = i
			}
		}
		sum = 0
	}
	return subs
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func main() {
	// intList := []int{1,2,-3,3,1,-4,2}
	intList := []int{6, -1, -3, 4, -2, 2, 4, 6, -12, -7}

	fmt.Println(zeroSumSubArr(intList))
}
