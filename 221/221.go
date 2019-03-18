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

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func decreaseSquareArea(height []int) int {
	maxLen := 0
	for i := range height {
		maxLen = getMax(maxLen, getMin(i+1, height[i]))
		if i+1 >= height[i] {
			break
		}
	}
	return maxLen * maxLen
}

func increaseSquareArea(height []int) int {
	maxLen := 0
	heightLen := len(height)
	for i := range height {
		maxLen = getMax(maxLen, getMin(heightLen-i, height[i]))
		if heightLen-i <= maxLen {
			break
		}
	}
	return maxLen * maxLen
}

func peakSquareArea(height []int) int {
	maxLen := 0
	i, j := 0, len(height)-1
	for i <= j {
		maxLen = getMax(maxLen, getMin(j-i+1, getMin(height[i], height[j])))
		if j-i+1 <= maxLen {
			break
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxLen * maxLen
}

func crossMinSquareArea(height []int, idx int) int {
	maxHeight := height[idx]
	left := idx - 1
	right := idx + 1
	for maxHeight > 0 {
		for left >= 0 && left > idx-maxHeight && height[left] >= maxHeight {
			left--
		}
		if idx-left >= maxHeight {
			return maxHeight * maxHeight
		}
		for right < len(height) && right < idx+maxHeight && height[right] >= maxHeight {
			right++
		}
		if right-left-1 >= maxHeight {
			return maxHeight * maxHeight
		}
		if left >= 0 && height[left] < maxHeight && right < len(height) && height[right] < maxHeight {
			maxHeight = getMax(height[left], height[right])
		} else if left >= 0 && height[left] < maxHeight {
			maxHeight = height[left]
		} else if right < len(height) && height[right] < maxHeight {
			maxHeight = height[right]
		} else {
			break
		}
	}
	return maxHeight * maxHeight
}

func maxSquareArea(height []int) int {
	left := 0
	lefti := left + 1
	heightLen := len(height)
	//skip begining decrease
	for lefti < heightLen && height[lefti] <= height[lefti-1] {
		lefti++
	}
	res := 0
	if lefti > left+1 {
		res = decreaseSquareArea(height[:lefti])
		left = lefti - 1
		res = getMax(res, crossMinSquareArea(height, left))
		if lefti < heightLen {
			for left > 0 && height[left] == height[left-1] {
				left--
			}
		}
	}
	for left < heightLen {
		lefti = left + 1
		for lefti < heightLen && height[lefti] >= height[lefti-1] {
			lefti++
		}
		if lefti == heightLen {
			res = getMax(res, increaseSquareArea(height[left:]))
			left = lefti
		} else {
			for lefti < heightLen && height[lefti] <= height[lefti-1] {
				lefti++
			}
			res = getMax(res, peakSquareArea(height[left:lefti]))
			left = lefti - 1
			res = getMax(res, crossMinSquareArea(height, left))
			if lefti < heightLen {
				for left > 0 && height[left] == height[left-1] {
					left--
				}
			}
		}
	}
	return res
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	cols := len(matrix[0])
	height := make([]int, cols)
	res := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		temp := maxSquareArea(height)
		if temp > res {
			res = temp
		}
	}
	return res
}

func main() {
	matrix := [][]byte{{'0', '1'}}
	fmt.Println(maximalSquare(matrix))
}
