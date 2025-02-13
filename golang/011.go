package main

func insertionSort(arr []int) []int {
	for i := range arr {
		k := i
		for k > 0 && arr[k-1] > arr[k] {
			arr[k-1], arr[k] = arr[k], arr[k-1]
			k--
		}

	}
	return arr
}

// func main(){
// 	// fmt.Println(insertionSort([]int{2,3,2,4,5,6,3,4,2,5}))
// 	fmt.Println(insertionSort([]int{4,2,5,6,6,23,23,3,42,6}))
// }
