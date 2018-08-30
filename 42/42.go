package main

import (
	"fmt"
)

func leftToRight(height []int) int {
	currentH := height[0]
	res := 0
	for i := 1; i < len(height); i++ {
		if height[i] >= currentH {
			currentH = height[i]
		} else {
			res += currentH - height[i]
		}
	}
	return res
}

func rightToLeft(height []int) int {
	currentH := height[len(height)-1]
	res := 0
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] >= currentH {
			currentH = height[i]
		} else {
			res += currentH - height[i]
		}
	}
	return res
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	left := height[0]
	right := height[len(height)-1]
	maxR := 0
	maxIdx := 0
	for i := 1; i < len(height)-1; i++ {
		if maxR < height[i] {
			maxR = height[i]
			maxIdx = i
		}
	}

	res := 0

	if maxR >= left && maxR >= right {
		res += leftToRight(height[:maxIdx+1]) + rightToLeft(height[maxIdx:])
	} else if left < right {
		res = leftToRight(height)
	} else {
		res = rightToLeft(height)
	}
	return res
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
