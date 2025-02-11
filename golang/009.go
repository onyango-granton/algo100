package main

import "fmt"

func getFirstAndLastOcc(arr []int,pointer,target, first, last int)[]int{
	if pointer == len(arr){
		return []int{first, last}
	}
	if arr[pointer] == target && first == 0{
		first = pointer
	} else if arr[pointer] == target && first != 0{
		last = pointer
	}
	return getFirstAndLastOcc(arr, pointer + 1, target,first, last)
}

func main(){
	fmt.Println(getFirstAndLastOcc([]int{2,3,4,5,6,7,8,8},0,8,0,0))
}