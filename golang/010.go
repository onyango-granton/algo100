package main

import "fmt"

func bubbleSort(arr []int)[]int{
	for range arr{
		for i := range arr{
		if i+1 < len(arr){
			if arr[i] > arr[i+1]{
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}}

	return arr
}

func recursiveCallBubbleSort(arr []int, pointer, iteration int) []int {
	//pointer reset
	if pointer == len(arr){
		pointer = 0
		iteration++
	}
	//base case
	if iteration == len(arr){
		return arr
	}
	if pointer + 1 < len(arr) && arr[pointer] > arr[pointer + 1]{
		arr[pointer], arr[pointer + 1] = arr[pointer + 1], arr[pointer]
	}
	return recursiveCallBubbleSort(arr, pointer+1, iteration)
}

func main(){
	fmt.Println(recursiveCallBubbleSort([]int{3,4,5,6,7,2,223,4,4,2},0,0))
}