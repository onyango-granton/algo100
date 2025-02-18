package main

import "fmt"

func findLargestSum(arr []int) int {
	// res := arr[0]
	iteration := 0
	for i :=1; i < len(arr); i++{
		if iteration+arr[i] < iteration{
			iteration = 0
		} else {
			iteration += arr[i]
		}
		fmt.Println(arr[i],iteration)
	}
	return iteration
}

func max(num1, num2 int) int {
	if num1 > num2{
		return num1
	} else {
		return num2
	}
}

func main() {
	arr1 := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(findLargestSum(arr1))
}