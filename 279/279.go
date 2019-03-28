package main

import (
	"fmt"
	"math"
)

func recSol(n int, sol map[int]int) int {
	sqrtn := int(math.Sqrt(float64(n)))
	if sqrtn*sqrtn == n {
		sol[n] = 1
		return 1
	}
	sol[n] = n
	for k := sqrtn; k > 0; k-- {
		remain := n - k*k
		if _, ok := sol[remain]; !ok {
			recSol(remain, sol)
		}
		v := sol[remain]
		if v+1 < sol[n] {
			sol[n] = v + 1
		}
	}
	return sol[n]
}

func numSquares(n int) int {
	sol := make(map[int]int)
	return recSol(n, sol)
}

func main() {
	fmt.Println(numSquares(12))
}
