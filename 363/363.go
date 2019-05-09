package main

import (
	"fmt"
)

func sumAll(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	trans := false
	if cols < rows {
		trans = true
	}
	var res [][]int
	if trans {
		res = make([][]int, cols+1)
		for i := range res {
			res[i] = make([]int, rows+1)
		}
	} else {
		res = make([][]int, rows+1)
		for i := range res {
			res[i] = make([]int, cols+1)
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if trans {
				res[j+1][i+1] = res[j][i+1] + matrix[i][j]
			} else {
				res[i+1][j+1] = res[i+1][j] + matrix[i][j]
			}
		}
		for j := 0; j < cols; j++ {
			if trans {
				res[j+1][i+1] += res[j+1][i]
			} else {
				res[i+1][j+1] += res[i][j+1]
			}
		}
	}
	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func halfFind(nums []int, target int) (int, bool) {
	if len(nums) == 0 {
		return -1, false
	}
	if len(nums) == 1 {
		if nums[0] < target {
			return -1, false
		}
		return 0, nums[0] == target
	}
	if len(nums) == 2 {
		if nums[0] < target {
			return -1, false
		} else if nums[1] < target {
			return 0, nums[0] == target
		}
		return 1, nums[1] == target
	}
	mid := len(nums) / 2
	if nums[mid] == target {
		return mid, true
	} else if nums[mid] < target {
		return halfFind(nums[:mid], target)
	} else {
		idx, found := halfFind(nums[mid+1:], target)
		return idx + mid + 1, found
	}
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	sums := sumAll(matrix)
	rows, cols := len(sums), len(sums[0])
	res := MinInt
	for i := 1; i < rows && res < k; i++ {
		for j := i; j < rows && res < k; j++ {
			sumNow := []int{0}
			for col := 1; col < cols && res < k; col++ {
				nowSum := sums[j][col] - sums[i-1][col]
				target := nowSum - k
				idx, _ := halfFind(sumNow, target)
				if idx >= 0 {
					res = getMax(res, nowSum-sumNow[idx])
				}
				if res < k {
					idx, found := halfFind(sumNow, nowSum)
					if !found {
						temp := append([]int{nowSum}, sumNow[idx+1:]...)
						if idx == -1 {
							sumNow = temp
						} else {
							sumNow = append(sumNow[:idx+1], temp...)
						}
					}
				}
			}
		}
	}
	return res
}

func main() {
	matrix0 := [][]int{{1, 0, 1}, {0, -2, 3}}
	k0 := 2
	fmt.Println(maxSumSubmatrix(matrix0, k0))

	matrix1 := [][]int{{2, 2, -1}}
	k1 := 0
	fmt.Println(maxSumSubmatrix(matrix1, k1))

	matrix2 := [][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}
	k2 := 3
	fmt.Println(maxSumSubmatrix(matrix2, k2))
}
