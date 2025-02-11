package main

import "fmt"

func findNum(arr []int, target int, start, stop int) int {
	if start  > stop {
		return -1
	}
	mid := start +(stop - start )/ 2

	// fmt.Println(mid)
	

	if arr[mid] == target{
		return mid
	}

	if target > arr[mid]{
		return findNum(arr, target,mid+1,stop)
	} else {
		return findNum(arr, target, start, mid-1 )
	}
}

func main(){
	fmt.Println(findNum([]int{2,3,5,6,7,8},10,0,5))
}