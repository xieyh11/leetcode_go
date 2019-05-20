package main

import (
	"fmt"
)

func canCross(stones []int) bool {
	if stones[1] != 1 {
		return false
	}
	res := make([]map[int]bool, len(stones))
	last := len(stones) - 1
	for i := len(stones) - 2; i >= 0; i-- {
		res[i] = make(map[int]bool)
		for j := i + 1; j < len(stones); j++ {
			if j != last {
				k := stones[j] - stones[i]
				if _, ok := res[j][k-1]; ok {
					res[i][k] = true
				} else if _, ok = res[j][k]; ok {
					res[i][k] = true
				} else if _, ok = res[j][k+1]; ok {
					res[i][k] = true
				}
			} else {
				res[i][stones[last]-stones[i]] = true
			}
		}
	}
	_, resOk := res[0][1]
	return resOk
}

func main() {
	stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
	fmt.Println(canCross(stones))
	stones = []int{0, 1, 2, 3, 4, 8, 9, 11}
	fmt.Println(canCross(stones))
}
