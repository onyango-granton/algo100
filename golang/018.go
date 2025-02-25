package main

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2

	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(leftArr, rightArr []int) []int {
	mainArr := make([]int, len(leftArr)+len(rightArr))
	lpointer, rpointer, mpointer := 0, 0, 0

	for lpointer < len(leftArr) && rpointer < len(rightArr) {
		if leftArr[lpointer] < rightArr[rpointer] {
			mainArr[mpointer] = leftArr[lpointer]
			lpointer++
		} else {
			mainArr[mpointer] = rightArr[rpointer]
			rpointer++
		}
		mpointer++
	}

	for lpointer < len(leftArr) {
		mainArr[mpointer] = leftArr[lpointer]
		mpointer++
		lpointer++
	}

	for rpointer < len(rightArr) {
		mainArr[mpointer] = rightArr[rpointer]
		mpointer++
		rpointer++
	}

	return mainArr
}

func twoSum(arr []int, target int) [][]int {
	arr = mergeSort(arr)
	indexes := [][]int{}

	intMap := make(map[int]bool)

	for i := range arr {
		intMap[i] = false
	}

	lpointer, rpointer := 0, len(arr)-1

	for lpointer < rpointer {
		if arr[lpointer]+arr[rpointer] == target {
			if !intMap[lpointer] && !intMap[rpointer] {
				indexes = append(indexes, []int{lpointer, rpointer})
				intMap[lpointer] = true
				intMap[rpointer] = true
			}

			lpointer++
			rpointer--
		} else if arr[lpointer]+arr[rpointer] > target {
			if arr[lpointer] > target && lpointer-1 >= 0 {
				lpointer--
			} else {
				rpointer--
			}
		} else if arr[lpointer]+arr[rpointer] < target {
			if arr[rpointer] < target && rpointer+1 < len(arr) {
				rpointer++
			} else {
				lpointer++
			}
		}
	}

	return indexes
}

// func main() {
// 	arr1 := []int{2, 3, 2, 5, 7, 8, 5, 9, 6}
// 	fmt.Println(mergeSort(arr1))
// 	fmt.Println(twoSum(arr1,9))
// }
