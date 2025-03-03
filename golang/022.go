package main

import "fmt"

// 1. Write a normal recursive function
// 2. Create a map function for x imma get y
// 3. Store x and y to the map
// 4. return y for each time you have x

var memo = make(map[int]int)

func fibonacci(position int) int {
	// base case
	if position < 2{
		return position
	}
	// memo case
	if val, found := memo[position]; found{
		return val
	}
	// mapping values
	memo[position] = fibonacci(position - 1) + fibonacci(position - 2)

	// return value for each loop
	return memo[position]
}

func main(){
	fmt.Println(fibonacci(6))
}