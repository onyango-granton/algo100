package main

import "fmt"

func tower_of_hanoi(n int, source, auxiliary, destination string){
	if n == 0{
		return
	}
	fmt.Println(n)
	tower_of_hanoi(n-1,source, destination, auxiliary)
	fmt.Println("Moved",n,"from",source,"to",destination)
	tower_of_hanoi(n-1,auxiliary,source,destination)
}

func main() {
	tower_of_hanoi(3,"A","B","C")
}