package main

import "fmt"

type Route struct {
	From City
	To   City
	Dist int
	Shortest bool
}

type City struct {
	Name string
	Visited bool
}

func getRoutes(distMatrix [][]int) {
	listRoutes := []Route{}
	citiesList := []string{"A", "B", "C", "D"}
	routes := make(map[string]bool)
	shortestDist := distMatrix[0][1]
	for i := 0; i < len(distMatrix); i++ {
		for j := 0; j < len(distMatrix[i]); j++ {
			if citiesList[i] > citiesList[j]{
				if citiesList[j] != citiesList[i] && !routes[citiesList[j]+citiesList[i]] {
					newRoute := Route{From: City{Name: citiesList[j]}, To: City{Name: citiesList[i]}, Dist: distMatrix[i][j]}
					if distMatrix[i][j] <= shortestDist{
						newRoute.Shortest = true
					}
					listRoutes = append(listRoutes, newRoute)
					routes[citiesList[j]+citiesList[i]] = true
				} else {
					continue
				}
				
			} else {
				if citiesList[j] != citiesList[i] && !routes[citiesList[i]+citiesList[j]]{
					newRoute := Route{From: City{Name: citiesList[i]}, To: City{Name: citiesList[j]}, Dist: distMatrix[i][j]}
					if distMatrix[i][j] <= shortestDist{
						newRoute.Shortest = true
					}
					listRoutes = append(listRoutes, newRoute)
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

func findRoute(routeList []Route){

}

func main() {
	distMatrix := [][]int{
		{0,10,15,20},
		{10,0,35,25},
		{15,35,0,5},
		{20,25,5,0},
	}
	getRoutes(distMatrix)
}