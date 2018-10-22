package main

import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	if len(heights) == 2 {
		min := heights[0]
		max := heights[1]
		if min > heights[1] {
			max = min
			min = heights[1]
		}
		if max > 2*min {
			return max
		} else {
			return 2 * min
		}
	}

	min := heights[0]
	minI := []int{0}
	for i := 1; i < len(heights); i++ {
		if heights[i] == min {
			minI = append(minI, i)
		} else if heights[i] < min {
			min = heights[i]
			minI = []int{i}
		}
	}
	current := min * len(heights)
	temp := largestRectangleArea(heights[:minI[0]])
	if temp > current {
		current = temp
	}
	temp = largestRectangleArea(heights[minI[len(minI)-1]+1:])
	if temp > current {
		current = temp
	}

	for i := 1; i < len(minI); i++ {
		temp = largestRectangleArea(heights[(minI[i-1] + 1):minI[i]])
		if temp > current {
			current = temp
		}
	}

	return current
}

func main() {
	heights := []int{9, 0}
	fmt.Println(largestRectangleArea(heights))
}
