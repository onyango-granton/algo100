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
	leftP, rightP, mainP := 0, 0, 0

	for leftP < len(leftArr) && rightP < len(rightArr) {
		if leftArr[leftP] <= rightArr[rightP] {
			mainArr[mainP] = leftArr[leftP]
			leftP++
		} else {
			mainArr[mainP] = rightArr[rightP]
			rightP++
		}
		mainP++
	}

	for leftP < len(leftArr) {
		mainArr[mainP] = leftArr[leftP]
		leftP++
		mainP++
	}

	for rightP < len(rightArr) {
		mainArr[mainP] = rightArr[rightP]
		rightP++
		mainP++
	}

	return mainArr

}

func findLongestConsecutiveSequence(arr []int) int {
	sortedArr := mergeSort(arr)

	listArr := [][]int{}
	var max int

	indexArr := []int{sortedArr[0]}

	for i := 1; i < len(sortedArr); i++ {
		if sortedArr[i]-indexArr[len(indexArr)-1] == 1 {
			indexArr = append(indexArr, sortedArr[i])
		} else {
			listArr = append(listArr, indexArr)
			indexArr = []int{sortedArr[i]}
		}
		if i == len(sortedArr)-1 {
			listArr = append(listArr, indexArr)
		}
	}

	for _, ch := range listArr {
		if len(ch) > max {
			max = len(ch)
		}
	}

	return max
}

// func main(){
// 	arr1 := []int{0,3,7,2,5,8,4,6,0,1}
// 	fmt.Println(findLongestConsecutiveSequence(arr1))
// }
