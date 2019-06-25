package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gScan := len(g) - 1
	sScan := len(s) - 1
	res := 0
	for gScan >= 0 && sScan >= 0 {
		if g[gScan] <= s[sScan] {
			res++
			sScan--
		}
		gScan--
	}
	return res
}

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	fmt.Println(findContentChildren(g, s))

	g = []int{1, 2, 3}
	s = []int{1, 2}
	fmt.Println(findContentChildren(g, s))
}
