package main

import "fmt"

func zeroSumSubArr(arr []int) []int{
	intMap := make(map[int]int)

	sum := 0
	for i,ch := range arr{
		sum += ch
		toZero := sum - 0

		val, ok := intMap[toZero]
		if ok {
			return []int{val, i}
		}
		intMap[ch] = i
	}

	return []int{}
}

func main(){
	// intList := []int{1,2,-3,3,1,-4,2}
	intList := []int{6, -1, -3, 4, -2, 2, 4, 6, -12, -7}

	start := 0
	stop := start

	subArrList := [][]int{}

	sum := 0

	for stop < len(intList){
		if sum + intList[stop] > sum && start > 0{
			start = stop
			sum = 0
		} else {
			sum += intList[stop]
			if sum == 0 {
				subArrList = append(subArrList, []int{start,stop})
				start = stop + 1	
			}
		}
		
		stop ++
	}

	fmt.Println(subArrList)
}