package main

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