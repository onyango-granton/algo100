package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) <= 1{
		return arr
	}
	pivot := arr[len(arr)-1]

	left := []int{}
	right := []int{}

	for _, ch := range arr[:len(arr)-1]{
		if ch <= pivot{
			left = append(left, ch)
		}
		if ch > pivot{
			right = append(right, ch)
		}
	}

	return append(append(quicksort(left),pivot),quicksort(right)...)
}


func main(){
	fmt.Println(quicksort([]int{3,6,8,3,6,7,6,4,2,6,8,0}))
}