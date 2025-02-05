package main

import "fmt"

func removeDuplicates(arr []int) []int{
	start:= 0
	stop := len(arr) - 1
	pointer := arr[0]
	for start <= stop{
		if start == 0{
			arr = append(arr, pointer)
		}
		if arr[start] != pointer{
			arr = append(arr, arr[start])
			pointer = arr[start]
		}
		start++
	}
	return arr[stop+1:]
}

func main() {
	fmt.Println(removeDuplicates([]int{3,3,3}))
}