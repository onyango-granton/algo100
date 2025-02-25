package main

func findPairWithSum(arr []int, target int) []int {
	intMap := make(map[int]int)
	intList := []int{}
	indexList := []int{}

	for i, ch := range arr {
		intMap[ch] = target - arr[i]
	}

	for k, v := range intMap {
		if intMap[v] != 0 {
			intList = append(intList, []int{k, v}...)
			break
		}
	}

	for i := 0; i < 2; {
		for k := 0; k < len(arr); k++ {
			if intList[i] == arr[k] {
				indexList = append(indexList, k)
				i++
			}
			if i == 2 {
				break
			}
		}
	}

	return indexList
}

// func main() {
// 	arr := []int{2,3,4,5,6,6}
// 	target := 9
// 	// intMap := make(map[int]int)

// 	fmt.Println(findPairWithSum(arr, target))
// }
