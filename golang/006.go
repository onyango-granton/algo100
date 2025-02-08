package main

func moveToLast(arr []int, index, nonZeroPointer int) []int {
	if index >= len(arr) {
		arr[nonZeroPointer] = 0
		nonZeroPointer++
		if nonZeroPointer == len(arr) {
			return arr
		}
	} else if arr[index] != 0 {
		arr[nonZeroPointer] = arr[index]
		nonZeroPointer++
	}

	return moveToLast(arr, index+1, nonZeroPointer)
}
