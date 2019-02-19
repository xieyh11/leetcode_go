package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return -1
	}
	startIdx := 0
	endIdx := 0
	contain := 0
	currentVal := gas[0] - cost[0]
	for !(startIdx == endIdx && contain > 0) {
		if currentVal >= 0 {
			endIdx++
			if endIdx >= len(gas) {
				endIdx -= len(gas)
			}
			currentVal += gas[endIdx] - cost[endIdx]
		} else {
			startIdx--
			if startIdx < 0 {
				startIdx += len(gas)
			}
			currentVal += gas[startIdx] - cost[startIdx]
		}
		contain++
	}
	if currentVal >= 0 {
		return startIdx
	} else {
		return -1
	}
}

func main() {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}
