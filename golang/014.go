package main

import "fmt"

func getUnion(arr1, arr2 []int) []int {
	numMap := make(map[int]bool)
	// resArr := []int{}
	var resArr []int

	// append creates a new array in memory
	// for _,ch := range append(arr1,arr2...){
	// 	if numMap[ch] == false{
	// 		resArr = append(resArr, ch)
	// 		numMap[ch] = true
	// 	}
	// }

	for _,num := range arr1{ //num for readability
		if !numMap[num] { //readability
			resArr = append(resArr, num)
			numMap[num] = true
		}
	}

	for _,num := range arr2{
		if !numMap[num] {
			resArr = append(resArr, num)
			numMap[num] = true
		}
	}

	return resArr
}

func main(){
	arr1 := []int{2,3,5,6,7,4,5,6,2,5} 
	arr2 := []int{3,4,5,1,4,5,3,5,3}
	fmt.Println(getUnion(arr1,arr2))
}