package main

import (
	"fmt"
)

type Bucket struct {
	Used bool
	Max  int
	Min  int
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func bucketSolver(nums []int) int {
	numOf := len(nums)
	if numOf < 2 {
		return 0
	}
	min, max := nums[0], nums[0]
	for _, num := range nums {
		min = getMin(min, num)
		max = getMax(max, num)
	}
	bucketSize := getMax(1, (max-min)/(numOf-1))
	bucketNum := (max-min)/bucketSize + 1
	buckets := make([]Bucket, bucketNum)

	for _, num := range nums {
		bucketIdx := (num - min) / bucketSize
		if !buckets[bucketIdx].Used {
			buckets[bucketIdx].Min = num
			buckets[bucketIdx].Max = num
		} else {
			buckets[bucketIdx].Min = getMin(num, buckets[bucketIdx].Min)
			buckets[bucketIdx].Max = getMax(num, buckets[bucketIdx].Max)
		}
		buckets[bucketIdx].Used = true
	}

	maxGap := 0
	prevBucketMax := min
	for _, bucket := range buckets {
		if bucket.Used {
			maxGap = getMax(maxGap, bucket.Min-prevBucketMax)
			prevBucketMax = bucket.Max
		}
	}
	return maxGap
}

func radixSortSolver(nums []int) int {
	numOf := len(nums)
	if numOf < 2 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		max = getMax(num, max)
	}
	expMove := uint(0)
	aux := make([]int, numOf)
	for expMove < 32 && (max>>expMove) > 0 {
		count := []int{0, 0}
		for _, num := range nums {
			count[(num>>expMove)&1]++
		}
		count[1] += count[0]
		for i := numOf - 1; i >= 0; i-- {
			at := (nums[i] >> expMove) & 1
			count[at]--
			aux[count[at]] = nums[i]
		}
		copy(nums, aux)
		expMove++
	}
	maxGap := 0
	for i := 1; i < numOf; i++ {
		maxGap = getMax(maxGap, nums[i]-nums[i-1])
	}
	return maxGap
}

func maximumGap(nums []int) int {
	return radixSortSolver(nums)
}

func main() {
	nums := []int{3, 6, 9, 1}
	fmt.Println(maximumGap(nums))
}
