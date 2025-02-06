package main

import "fmt"

func divide(arr []int) []int {
	if len(arr) < 2{
		return arr
	}

	midPoint := len(arr) / 2

	leftArr := divide(arr[:midPoint])
	rightArr := divide(arr[midPoint:])

	return conquer(leftArr, rightArr)
}

func conquer(leftArr, rightArr []int) []int{
	mainArr := make([]int, len(leftArr)+len(rightArr))
	mainP, rightP, leftP := 0,0,0

	for rightP < len(rightArr) && leftP < len(leftArr){
		if rightArr[rightP] <= leftArr[leftP]{
			mainArr[mainP] = rightArr[rightP]
			rightP++
		} else {
			mainArr[mainP] = leftArr[leftP]
			leftP++
		}
		mainP++
	}

	for leftP < len(leftArr){
		mainArr[mainP] = leftArr[leftP]
		mainP++
		leftP++
	}

	for rightP < len(rightArr){
		mainArr[mainP] = rightArr[rightP]
		mainP++
		rightP++
	}

	return mainArr

}

func main(){
	fmt.Println(divide([]int{4,6,2,2,4,2,56,8,3,23}))
}