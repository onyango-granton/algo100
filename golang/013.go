package main

import "fmt"

func findMissingNum(arr []int)int{
	max, min := arr[0], arr[0]
	missingNum := -1
	for _,ch := range arr{
		if ch > max{
			max = ch
		}
		if ch < min{
			min = ch
		}
	}
	//cover for the fact that indexing starts at 0 and max-min returns one less
	array := make([]int, max-min+2)
	fmt.Println(array)

	for _,ch := range arr {
		array[ch] = 1
	}

	for i:=1; i < len(array); i++{
		if array[i] == 0{
			missingNum = i
		}
	}

	return missingNum
}

func main(){
	fmt.Println(findMissingNum([]int{3, 7, 1, 2, 8, 4, 5}))
}