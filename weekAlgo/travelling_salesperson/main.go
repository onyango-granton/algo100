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
	routes := make(map[string]bool)
	for i := 0; i < len(distMatrix); i++ {
		for j := 0; j < len(distMatrix[i]); j++ {
			if citiesList[i] > citiesList[j]{
				if citiesList[j] != citiesList[i] && !routes[citiesList[j]+citiesList[i]] {
					listRoutes = append(listRoutes, Route{From: citiesList[j], To: citiesList[i], Dist: distMatrix[i][j]})
					routes[citiesList[j]+citiesList[i]] = true
				} else {
					continue
				}
				
			} else {
				if citiesList[j] != citiesList[i] && !routes[citiesList[i]+citiesList[j]]{
					listRoutes = append(listRoutes, Route{From: citiesList[i], To: citiesList[j], Dist: distMatrix[i][j]})
					routes[citiesList[i]+citiesList[j]] = true
				} else {
					continue
				}
				//listRoutes = append(listRoutes, Route{From: citiesList[i], To: citiesList[j], Dist: distMatrix[i][j]})
			}
			
		}
	}
	for _, ch := range listRoutes {
		fmt.Println(ch)
	}
}

func main() {
	distMatrix := [][]int{
		{0,10,15,20},
		{10,0,35,25},
		{15,35,0,30},
		{20,25,30,0},
	}
	getRoutes(distMatrix)
}