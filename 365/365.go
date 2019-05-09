package main

import (
	"fmt"
)

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Pair [2]int

func recSol(x, y int, xC, yC int, reachable map[Pair]bool, searched map[Pair]bool) bool {
	pair := Pair{x, y}
	if v, ok := reachable[pair]; ok {
		return v
	}
	res := false
	searchPair := Pair{xC, y}
	if v, ok := searched[searchPair]; x < xC && (!ok || !v) {
		searched[searchPair] = true
		res = recSol(xC, y, xC, yC, reachable, searched)
		searched[searchPair] = false
	}
	searchPair = Pair{x, yC}
	if v, ok := searched[searchPair]; !res && y < yC && (!ok || !v) {
		searched[searchPair] = true
		res = recSol(x, yC, xC, yC, reachable, searched)
		searched[searchPair] = false
	}

	searchPair = Pair{0, y}
	if v, ok := searched[searchPair]; !res && x > 0 && (!ok || !v) {
		searched[searchPair] = true
		res = recSol(0, y, xC, yC, reachable, searched)
		searched[searchPair] = false
	}
	searchPair = Pair{x, 0}
	if v, ok := searched[searchPair]; !res && y > 0 && (!ok || !v) {
		searched[searchPair] = true
		res = recSol(x, 0, xC, yC, reachable, searched)
		searched[searchPair] = false
	}

	if !res && x > 0 && y < yC {
		trans := getMin(x, yC-y)
		searchPair = Pair{x - trans, y + trans}
		if v, ok := searched[searchPair]; !ok || !v {
			searched[searchPair] = true
			res = recSol(x-trans, y+trans, xC, yC, reachable, searched)
			searched[searchPair] = false
		}
	}
	if !res && x < xC && y > 0 {
		trans := getMin(xC-x, y)
		searchPair = Pair{x + trans, y - trans}
		if v, ok := searched[searchPair]; !ok || !v {
			searched[searchPair] = true
			res = recSol(x+trans, y-trans, xC, yC, reachable, searched)
			searched[searchPair] = false
		}
	}
	reachable[pair] = res
	return res
}

func canMeasureWater(x int, y int, z int) bool {
	res := make(map[Pair]bool)
	searched := make(map[Pair]bool)
	for i := 0; i <= z && i <= x; i++ {
		if z-i <= y {
			res[Pair{i, z - i}] = true
		}
	}
	searched[Pair{0, 0}] = true
	return recSol(0, 0, x, y, res, searched)
}

func main() {
	fmt.Println(canMeasureWater(3, 5, 4))

	fmt.Println(canMeasureWater(2, 6, 5))

	fmt.Println(canMeasureWater(22003, 31237, 1))
}
