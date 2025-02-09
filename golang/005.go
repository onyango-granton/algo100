package main

func divide(arr []int) []int {
	//base case
	if len(arr) < 2 {
		return arr
	}

	midpoint := len(arr) / 2
	leftArr := divide(arr[:midpoint])
	rightArr := divide(arr[midpoint:])

	return conquer(leftArr, rightArr)
}

func conquer(leftArr, rightArr []int) []int {
	mainArr := make([]int, len(leftArr)+len(rightArr))
	mainP, leftP, rightP := 0, 0, 0

	for leftP < len(leftArr) && rightP < len(rightArr) {
		if leftArr[leftP] < rightArr[rightP] {
			mainArr[mainP] = leftArr[leftP]
			mainP++
			leftP++
		} else {
			mainArr[mainP] = rightArr[rightP]
			mainP++
			rightP++
		}

	}

	for leftP < len(leftArr) {
		mainArr[mainP] = leftArr[leftP]
		mainP++
		leftP++
	}

	for rightP < len(rightArr) {
		mainArr[mainP] = rightArr[rightP]
		mainP++
		rightP++
	}

	return mainArr
}

func compareArr(arrOne, arrTwo []int) []int {
	sortedArrOne, sortedArrTwo := divide(arrOne), divide(arrTwo)
	pointerOne, pointerTwo := 0, 0
	resArr := []int{}

	if len(sortedArrTwo) < len(sortedArrOne) {
		copyArr := sortedArrTwo
		sortedArrTwo = sortedArrOne
		sortedArrOne = copyArr
	}

	for pointerOne < len(sortedArrOne) && pointerTwo < len(sortedArrTwo) {
		if sortedArrOne[pointerOne] == sortedArrTwo[pointerTwo] {
			if len(resArr)-1 < 0 {
				resArr = append(resArr, sortedArrOne[pointerOne])
				pointerOne++
				pointerTwo++
			} else if sortedArrOne[pointerOne] != resArr[len(resArr)-1] {
				resArr = append(resArr, sortedArrOne[pointerOne])
				pointerOne++
				pointerTwo++
			} else {
				pointerOne++
				pointerTwo++
			}

		} else if sortedArrOne[pointerOne] > sortedArrTwo[pointerTwo] {
			pointerTwo++
		} else if sortedArrOne[pointerOne] < sortedArrTwo[pointerTwo] {
			pointerOne++
		}
	}

	return resArr
}

/*func main() {
	nums1 := []int{4, 5, 6, 7, 3, 23, 5}
	nums2 := []int{2, 5, 6, 2, 4, 4, 2, 8}
	// fmt.Println(divide([]int{3,6,8,3,24,5,5,3,42,4,67,6}))
	fmt.Println(compareArr(nums2, nums1))
}*/
