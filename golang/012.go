package main

import "fmt"

func mergeArr(arr1, arr2[]int) []int{
	mainArr := make([]int, len(arr1)+len(arr2))
	mainPointer := 0

	arr1Pointer := 0
	arr2Pointer := 0

	for arr1Pointer < len(arr1) && arr2Pointer < len(arr2){
		if arr1[arr1Pointer] < arr2[arr2Pointer]{
			mainArr[mainPointer] = arr1[arr1Pointer]
			arr1Pointer++
		} else {
			mainArr[mainPointer] = arr2[arr2Pointer]
			arr2Pointer++
		}
		mainPointer++
	}
	for arr1Pointer < len(arr1){

		mainArr[mainPointer] = arr1[arr1Pointer]
		arr1Pointer++
		mainPointer++
	}
	for arr2Pointer < len(arr2){
		mainArr[mainPointer] = arr2[arr2Pointer]
		arr2Pointer++
		mainPointer++
	}

	return mainArr
}

func main(){
	fmt.Println(mergeArr([]int{0, 7, 8},[]int{1,3,6}))
}