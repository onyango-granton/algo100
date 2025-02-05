package main

import (
	"unicode"
)

func isPalindrome(s string) bool {
	leftPointer := 0
	rightPointer := len(s) - 1

	for leftPointer <= rightPointer {
		if !isAlphanumeric(s[leftPointer]) {
			leftPointer++
			continue
		}
		if !isAlphanumeric(s[rightPointer]) {
			rightPointer--
			continue
		}
		if !isSame(s[leftPointer], s[rightPointer]) {
			return false
		}
		leftPointer++
		rightPointer--
	}
	return true
}

func isSame(char1, char2 byte) bool {
	return unicode.ToLower(rune(char1)) == unicode.ToLower(rune(char2))
}

func isAlphanumeric(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

// func main(){
// 	words := []string{"hello","de", "helo","0racc...ar0."}
// 	for _, word := range words{
// 		fmt.Println(isPalindrome(word))
// 	}
// }
