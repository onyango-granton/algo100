package main

import "fmt"

func findLargestSum(arr []int) int {
	maxSum := arr[0]
	sum := arr[0]

	for i:=1; i < len(arr);i++{
		if sum + arr[i] < arr[i]{
			sum = arr[i]
		} else {
			sum = sum + arr[i]
		}
		if maxSum < sum{
			maxSum = sum
		}
	}

	return maxSum
}

func main() {
	fmt.Println(findLargestSum([]int{5,4,-1,7,8}))
}