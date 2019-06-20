package main

import (
	"fmt"
)

func powDis(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

func numberOfBoomerangs(points [][]int) int {
	if len(points) < 3 {
		return 0
	}
	lengthNums := make([]map[int]int, len(points))
	for i := range lengthNums {
		lengthNums[i] = make(map[int]int)
	}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dis := powDis(points[i], points[j])
			lengthNums[i][dis]++
			lengthNums[j][dis]++
		}
	}
	res := 0
	for i := range lengthNums {
		for _, v := range lengthNums[i] {
			if v >= 2 {
				res += v * (v - 1)
			}
		}
	}
	return res
}

func main() {
	points := [][]int{{0, 0}, {1, 0}, {2, 0}}
	fmt.Println(numberOfBoomerangs(points))
}
