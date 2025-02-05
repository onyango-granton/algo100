package main

func findMaxMin(arr []int) []int {
	max, min := arr[0], arr[0]
	for _, ch := range arr {
		if ch > max {
			max = ch
		}
		if ch < min {
			min = ch
		}
	}
	return []int{max, min}
}

// func main() {
// 	fmt.Println(findMaxMin([]int{1, 4, 7, 2, 3, 4, 3, 6, 8}))
// }
