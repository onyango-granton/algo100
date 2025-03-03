package main

import "fmt"

func fibonacci(position int) int {
	if position < 2{
		return position
	}
	return fibonacci(position - 1) + fibonacci(position - 2)
}

func main(){
	fmt.Println(fibonacci(6))
}