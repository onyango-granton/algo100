package main

import "fmt"

func twoSum(arr []int, target int) []int {
	//indices := []int{}
	//map current iteration index to (target - arr[i])
	diffMap := make(map[int]int)

	for i, num := range arr {
		diff := target - arr[i]

		val, found := diffMap[num]

		//checks for precence of current arr index as key in map
		if found {
			return []int{val, i}
		}

		diffMap[diff] = i
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
