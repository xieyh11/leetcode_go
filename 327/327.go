package main

import (
	"fmt"
)

func isIn(num, lower, upper int) bool {
	return num >= lower && num <= upper
}

func recSol(sums []int, lower int, upper int) (int, []int) {
	if len(sums) == 0 {
		return 0, []int{}
	}
	if len(sums) == 1 {
		if isIn(sums[0], lower, upper) {
			return 1, sums
		}
		return 0, sums
	}
	if len(sums) == 2 {
		count := 0
		if isIn(sums[0], lower, upper) {
			count++
		}
		if isIn(sums[1], lower, upper) {
			count++
		}
		if isIn(sums[1]-sums[0], lower, upper) {
			count++
		}
		if sums[0] > sums[1] {
			sums[0], sums[1] = sums[1], sums[0]
		}
		return count, sums
	}
	mid := len(sums) / 2
	left, leftSums := recSol(sums[:mid], lower, upper)
	right, rightSums := recSol(sums[mid:], lower, upper)
	count := 0
	leftLower := 0
	leftUpper := 0
	rightIdx := 0
	for rightIdx < len(rightSums) {
		tempLower := rightSums[rightIdx] - upper
		tempUpper := rightSums[rightIdx] - lower
		for leftLower < len(leftSums) && tempLower > leftSums[leftLower] {
			leftLower++
		}
		for leftUpper < len(leftSums) && tempUpper >= leftSums[leftUpper] {
			leftUpper++
		}
		if leftLower < leftUpper {
			count += leftUpper - leftLower
		}
		rightIdx++
	}

	newSums := make([]int, 0)
	i, j := 0, 0
	for i < len(leftSums) && j < len(rightSums) {
		if leftSums[i] < rightSums[j] {
			newSums = append(newSums, leftSums[i])
			i++
		} else {
			newSums = append(newSums, rightSums[j])
			j++
		}
	}
	if i < len(leftSums) {
		newSums = append(newSums, leftSums[i:]...)
	}
	if j < len(rightSums) {
		newSums = append(newSums, rightSums[j:]...)
	}
	return count + left + right, newSums
}

func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	res, _ := recSol(sums, lower, upper)
	return res
}

func main() {
	nums := []int{2147483647, -2147483648, -1, 0}
	lower := -1
	upper := 0
	fmt.Println(countRangeSum(nums, lower, upper))
}
