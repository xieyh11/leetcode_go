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

func increasingTriplet(nums []int) bool {
	length2 := 0
	length1 := 0
	lenghtHave := [2]bool{}
	for _, num := range nums {
		if !lenghtHave[0] {
			length1 = num
			lenghtHave[0] = true
		} else {
			if !lenghtHave[1] {
				if length1 < num {
					lenghtHave[1] = true
					length2 = num
				}
			} else {
				if num > length2 {
					return true
				} else if num > length1 {
					length2 = num
				}
			}
			length1 = getMin(length1, num)
		}
	}
	return false
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	fmt.Println(increasingTriplet(nums1))

	nums2 := []int{1, 5, 2, 1, 3}
	fmt.Println(increasingTriplet(nums2))

	nums3 := []int{5, 4, 3, 2, 1}
	fmt.Println(increasingTriplet(nums3))
}
