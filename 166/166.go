package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	neg := false
	if numerator < 0 {
		numerator = -numerator
		neg = !neg
	}
	if denominator < 0 {
		denominator = -denominator
		neg = !neg
	}
	remain := numerator % denominator
	beforePoint := numerator / denominator
	afterPoint := []int{}
	remainMap := make(map[int]int)
	start := -1
	for remain != 0 {
		if idx, ok := remainMap[remain]; ok {
			start = idx
			break
		} else {
			remainMap[remain] = len(afterPoint)
			remain *= 10
			afterPoint = append(afterPoint, remain/denominator)
			remain %= denominator
		}
	}
	res := strconv.FormatInt(int64(beforePoint), 10)
	if neg && (beforePoint != 0 || len(afterPoint) > 0) {
		res = "-" + res
	}
	if len(afterPoint) > 0 {
		res += "."
		if start == -1 {
			start = len(afterPoint)
		}
		for i := 0; i < start; i++ {
			res += strconv.FormatInt(int64(afterPoint[i]), 10)
		}
		if start != len(afterPoint) {
			res += "("
			for i := start; i < len(afterPoint); i++ {
				res += strconv.FormatInt(int64(afterPoint[i]), 10)
			}
			res += ")"
		}
	}
	return res
}

func main() {
	fmt.Println(fractionToDecimal(0, -12))
}
