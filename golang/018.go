package main

func twoSum(arr []int, target int){
	intMap := make(map[int]int)

	for i, num := range arr{
		diff := num - target
		if num > target{
			continue
		}
		intMap[num] = diff
	}

	for key, value := range intMap{
		
	}
}