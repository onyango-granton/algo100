package main

func findSubArrayZero(arr []int) [][]int {
	var sum int
	//initalize subarray array
	subArr := [][]int{}

	//hash map that will map list of indeces whose sum are same
	//think of it as a wave each time it two trougs or crest are same in amplitude
	//sums to zero
	zeroSum := make(map[int][]int)

	for i, num := range arr {
		sum += num

		vals, found := zeroSum[sum]
		//if key already exsists
		// itearte through the values obtaining sub arr with start point at values plus one
		// and endpoint at index + 1
		if found {
			for _, value := range vals {
				subArr = append(subArr, []int{value + 1, i + 1})
			}
		}

		zeroSum[sum] = append(zeroSum[sum], i)
	}

	return subArr
}

// for a practical env approach use this arr

//func main() {
//	fmt.Println(findSubArrayZero([]int{6, -1, -3, 4, -2, 2, 4, 6, -12, -7}))
//}
