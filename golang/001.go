package main

import "fmt"

func revArr(arr []int) []int {
	resArr := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		resArr = append(resArr, arr[i])
	}
	return resArr
}

func revArrTwoPointer(arr []int) []int {
	leftPointer := 0
	rightPointer := len(arr) - 1

	for leftPointer < rightPointer {
		arr[leftPointer], arr[rightPointer] = arr[rightPointer], arr[leftPointer]
		leftPointer++
		rightPointer--
	}

	return arr
}

func main() {
	fmt.Println(revArrTwoPointer([]int{3, 6, 7, 8, 2, 4}))
}
