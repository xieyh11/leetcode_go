package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	version1Split := strings.Split(version1, ".")
	version2Split := strings.Split(version2, ".")
	minLen, maxLen := len(version1Split), len(version1Split)
	version1Long := true
	if minLen > len(version2Split) {
		minLen = len(version2Split)
	}
	if maxLen < len(version2Split) {
		maxLen = len(version2Split)
		version1Long = false
	}
	for i := 0; i < minLen; i++ {
		num1, _ := strconv.ParseInt(version1Split[i], 10, 64)
		num2, _ := strconv.ParseInt(version2Split[i], 10, 64)
		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
	}
	for i := minLen; i < maxLen; i++ {
		if version1Long {
			num, _ := strconv.ParseInt(version1Split[i], 10, 64)
			if num != 0 {
				return 1
			}
		} else {
			num, _ := strconv.ParseInt(version2Split[i], 10, 64)
			if num != 0 {
				return -1
			}
		}
	}
	return 0
}

func main() {
	version1 := "1.0.1"
	version2 := "1"
	fmt.Println(compareVersion(version1, version2))
}
