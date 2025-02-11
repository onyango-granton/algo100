package main

import "fmt"

func rotateArr(arr []int, positions int) []int {
	realPointer := 0
	copyArr := make([]int, len(arr))
	copy(copyArr,arr)

	fmt.Println(copyArr)

	for i := positions +1; i < len(arr); i++ {
		fmt.Print("here")
		arr[realPointer] = arr[i]
		realPointer++
	}
	fmt.Println(arr)



	for i := 0; i <= positions; i++ {
		arr[realPointer] = copyArr[i]
		realPointer++
	}

	return arr
}

// func main() {
// 	fmt.Println(rotateArr([]int{1,2,3,4,5,6,7}, 3))
// }