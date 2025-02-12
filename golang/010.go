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

func main(){
	fmt.Println(bubbleSort([]int{3,4,5,6,7,2,223,4,4,2}))
}