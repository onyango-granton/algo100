// non repeat char should not be exsisting
package main

import "fmt"

func findFirstNonRepeating(word string) string {
	charIndexList := make(map[rune][]int)
	leastIndex := len(word)
	for i,ch := range word{
		charIndexList[ch] = append(charIndexList[ch], i)
	}

	for _, indexList := range charIndexList{
		if len(indexList) != 1 {
			continue
		}
		if indexList[0] < leastIndex{
			leastIndex = indexList[0]
		}
	}

	return string(word[leastIndex])
}

func main() {
	fmt.Println(findFirstNonRepeating("leetcode"))
}