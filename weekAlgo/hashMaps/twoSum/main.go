package main

import "fmt"

func twoSum(arr []int, target int) []int {
	diffMap := make(map[int]int)

	for i, num := range arr {
		val, found := diffMap[num]

		if found {
			return []int{val, i}
		}

		diffMap[target-num] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
