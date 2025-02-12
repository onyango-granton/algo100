package main

import "fmt"

func getFirstAndLastOcc(arr []int,target, first, last int)[]int{
	position := binarySearch(arr, target, 0, len(arr))
	positionCopy :=  position
	for arr[position] == target{
		fmt.Println("here")
		first = position
		position --
	}
	for positionCopy < len(arr) && arr[positionCopy] == target{
		last = positionCopy
		positionCopy++
	}
	return []int{first, last}
}

func binarySearch(arr []int, target, start, stop int) int {
	if start > stop{
		return -1
	}

	midVal := start + (stop - start) / 2
	if target == arr[midVal]{
		return midVal
	} else if target > arr[midVal]{
		return binarySearch(arr, target, midVal + 1, stop)
	} else {
		return binarySearch(arr, target, start, midVal - 1)
	}
}

func main(){
	fmt.Println(getFirstAndLastOcc([]int{2,3,4,5,6,7,8,8},8,0,0))
	fmt.Println(binarySearch([]int{2,3,4,5,6,7,8,9},10,0,7))
}