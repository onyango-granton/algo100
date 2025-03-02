// non repeat char should not be exsisting
package main

import "fmt"

func findFirstNonRepeating(word string) string {
	firstIndex := len(word)

	charOcurrence := make(map[rune]int)
	charFirstIndex := make(map[rune]int)

	for i, ch := range word{
		//calculate total number of occurence per word
		charOcurrence[ch]++
		if _, found := charFirstIndex[ch]; !found{
			charFirstIndex[ch] = i
		}
	}

	for _,ch := range word{
		if charOcurrence[ch] == 1 && charFirstIndex[ch] < firstIndex{
			firstIndex = charFirstIndex[ch]
		}
	}

	if firstIndex == len(word){
		return ""
	}

	return string(word[firstIndex])

}

func main() {
	fmt.Println(findFirstNonRepeating("leetcode"))
}