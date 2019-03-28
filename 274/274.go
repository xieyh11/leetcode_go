package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	for i := range citations {
		if citations[i] >= len(citations)-i {
			return len(citations) - i
		}
	}
	return 0
}

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(citations))
}
