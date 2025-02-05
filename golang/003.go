package main

import "fmt"

func isPalindrome(s string) bool {
	leftPointer := 0
	rightPointer := len(s)-1

	for leftPointer <= rightPointer{
		if !(s[leftPointer] >= 'a' && s[leftPointer] <= 'z' || s[leftPointer] >= 'A' && s[leftPointer] <= 'Z'){
			leftPointer++
			continue
		}
		if !(s[rightPointer] >= 'a' && s[rightPointer] <= 'z' || s[rightPointer] >= 'A' && s[rightPointer] <= 'Z'){
			rightPointer--
			continue
		}
		if !isSame(s[leftPointer], s[rightPointer]){
			return false
		}
		leftPointer++
		rightPointer--		
	}
	return true
}

func isSame(char1, char2 byte) bool {
	return char1 == char2
}

func main(){
	words := []string{"hello", "helo","raccar"}
	for _, word := range words{
		fmt.Println(isPalindrome(word))
	}
}