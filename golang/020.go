package main

import (
	"fmt"
)

func zeroSumSubArr(arr []int) [][]int {

	subs := [][]int{}


	for i := 0; i < len(arr); i++{
		start := i
		stop := i
		sum := 0 
		for j := i; j < len(arr);j++{
			
			sum += arr[j]
			//fmt.Println(sum)
			if sum == 0 {
				stop = j
				subs = append(subs, []int{start, stop+1})
				start = j
				sum = 0
			}
		}
	}
	return subs
}

func main() {
	intList := []int{6, -1, -3, 4, -2, 2, 4, 6, -12, -7}

	fmt.Println(zeroSumSubArr(intList))
}
