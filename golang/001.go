package main

import "fmt"

func revArr(arr []int) []int {
	resArr := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		resArr = append(resArr, arr[i])
	}
	return resArr
}

func main() {
	fmt.Println(revArr([]int{3, 6, 7, 8, 2, 4}))
}
