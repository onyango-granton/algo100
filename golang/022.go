package main

import "fmt"

func fibonacci(num int) int {
	if num == 0{
		return 1
	}
	if num == 1{
		return 1
	}
	return fibonacci(num - 1) + fibonacci(num - 2)
}

func main(){
	fmt.Println(fibonacci(5))
}