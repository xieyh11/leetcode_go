package main

import (
	"fmt"
)

func recSol(citations []int, larger int) int {
	if len(citations) == 0 {
		return larger
	}
	if len(citations) == 1 {
		if citations[0] >= 1+larger {
			return 1 + larger
		}
		return larger
	}
	if len(citations) == 2 {
		if citations[0] >= 2+larger {
			return 2 + larger
		}
		if citations[1] >= 1+larger {
			return 1 + larger
		}
		return 0 + larger
	}
	mid := len(citations) / 2
	if citations[mid] == len(citations)-mid+larger {
		return len(citations) - mid + larger
	} else if citations[mid] > len(citations)-mid+larger {
		return recSol(citations[:mid], len(citations)-mid+larger)
	} else {
		return recSol(citations[mid+1:], larger)
	}
}

func hIndex(citations []int) int {
	return recSol(citations, 0)
}

func main() {
	citations := []int{0, 1, 3, 5, 6, 7}
	fmt.Println(hIndex(citations))
}
