package main

import (
	"fmt"
	"sort"
)

func recSol(edges map[string][]string, used map[string][]bool, nowAt string, length int) []string {
	v := edges[nowAt]
	vUsed := used[nowAt]
	for i := range v {
		if !vUsed[i] {
			vUsed[i] = true
			temp := recSol(edges, used, v[i], length-1)
			if len(temp) > 0 {
				return append([]string{nowAt}, temp...)
			}
			vUsed[i] = false
		}
	}
	if length == 1 {
		return []string{nowAt}
	}
	return []string{}
}

func findItinerary(tickets [][]string) []string {
	destination := make(map[string][]string)
	for _, pairs := range tickets {
		destination[pairs[0]] = append(destination[pairs[0]], pairs[1])
	}
	for _, v := range destination {
		sort.Strings(v)
	}

	used := make(map[string][]bool)
	for k, v := range destination {
		used[k] = make([]bool, len(v))
	}

	return recSol(destination, used, "JFK", len(tickets)+1)
}

func main() {
	str1 := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	fmt.Println(findItinerary(str1))

	str2 := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	fmt.Println(findItinerary(str2))

	str3 := [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}
	fmt.Println(findItinerary(str3))
}
