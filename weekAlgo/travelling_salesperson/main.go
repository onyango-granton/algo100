package main

import "fmt"

type Route struct {
	From string
	To   string
	Dist int
}

func getRoutes(distMatrix [][]int) {
	listRoutes := []Route{}
	citiesList := []string{"A", "B", "C", "D"}
	for i := 0; i < len(distMatrix); i++ {
		for j := 0; j < len(distMatrix[i]); j++ {
			listRoutes = append(listRoutes, Route{From: citiesList[i], To: citiesList[j], Dist: distMatrix[i][j]})
		}
	}
	for _, ch := range listRoutes {
		fmt.Println(ch)
	}
}

func main() {
	distMatrix := [][]int{
		{},
		{},
		{},
		{},
	}
}