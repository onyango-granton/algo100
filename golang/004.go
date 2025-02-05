package main

import "fmt"

func removeDuplicates(arr []int) []int{
	seen := make(map[int]bool)
	result := &[]int{}

	for _,num := range arr{
		if !seen[num]{
			*result = append(*result, num)
			seen[num] = true
		}
	}

	return *result
}

func main() {
	fmt.Println(removeDuplicates([]int{3,3,3}))
}