package main

import "fmt"

func findSubArrayZero(arr []int) [][]int{
	var sum int
	subArr := [][]int{}
	zeroSum := make(map[int][]int)
	for i, num := range arr {
		sum += num

		vals, found := zeroSum[sum]

		if found {
			for _, value := range vals {
				subArr = append(subArr, []int{value + 1, i + 1})
			}
		}

		zeroSum[sum] = append(zeroSum[sum], i)
	}

	return subArr
}

func main() {
	fmt.Println(findSubArrayZero([]int{6, -1, -3, 4, -2, 2, 4, 6, -12, -7}))
}