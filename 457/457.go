package main

import (
	"fmt"
)

func step(index, k, n int) int {
	index += k
	if index < 0 {
		return index%n + n
	}
	return index % n
}

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}
	for i := range nums {
		if nums[i] != 0 {
			direction := nums[i] > 0
			stepOne := step(i, nums[i], n)
			if (nums[stepOne] > 0) != direction {
				nums[i] = 0
				continue
			}
			stepTwo := step(stepOne, nums[stepOne], n)
			if (nums[stepTwo] > 0) != direction {
				nums[i] = 0
				nums[stepOne] = 0
				continue
			}
			for stepOne != stepTwo {
				stepOne = step(stepOne, nums[stepOne], n)
				stepTwo = step(stepTwo, nums[stepTwo], n)
				if (nums[stepTwo] > 0) != direction {
					break
				}
				stepTwo = step(stepTwo, nums[stepTwo], n)
				if (nums[stepTwo] > 0) != direction {
					break
				}
			}
			if (nums[stepTwo] > 0) != direction {
				stepOne = i
				for stepOne != stepTwo {
					stepOne, nums[stepOne] = step(stepOne, nums[stepOne], n), 0
				}
				continue
			}
			stepOne = i
			for stepOne != stepTwo {
				stepOne = step(stepOne, nums[stepOne], n)
				stepTwo = step(stepTwo, nums[stepTwo], n)
			}
			cycleScan := stepOne
			cycleLen := 1
			cycleScan, nums[cycleScan] = step(cycleScan, nums[cycleScan], n), 0
			for cycleScan != stepOne {
				cycleLen++
				cycleScan, nums[cycleScan] = step(cycleScan, nums[cycleScan], n), 0
			}
			if cycleLen > 1 {
				return true
			}
			cycleStart := stepOne
			stepOne = i
			for stepOne != cycleStart {
				stepOne, nums[stepOne] = step(stepOne, nums[stepOne], n), 0
			}
		}
	}
	return false
}

func main() {
	nums := []int{2, -1, 1, 2, 2}
	fmt.Println(circularArrayLoop(nums))

	nums = []int{-1, 2}
	fmt.Println(circularArrayLoop(nums))

	nums = []int{-2, 1, -1, -2, -2}
	fmt.Println(circularArrayLoop(nums))

	nums = []int{-1}
	fmt.Println(circularArrayLoop(nums))
}
