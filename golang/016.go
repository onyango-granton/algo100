package main

import "fmt"

func firstNonRepeatingChar(word string) int {
	charMap := make(map[string][]int)
	min := len(word)

	for i, ch := range word{
		charMap[string(ch)] = append(charMap[string(ch)], i)
	}

	for _, value := range charMap{
		if len(value) == 1{
			if value[0] < min{
				min = value[0]
			}
		}
	}

	if min >= 0 && min <= len(word)-1{
		return min
	} else {
		return -1
	}

	// return min
}

func main(){
	fmt.Println(firstNonRepeatingChar("aabb"))
}