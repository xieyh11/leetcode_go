package main

import (
	"fmt"
	"strings"
)

var Letters = []byte("efghinorstuvwxz")

var DigitCount = [][]int{{1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0},
	{2, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0},
	{1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0},
	{2, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0},
	{1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{1, 0, 0, 0, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

func countUpdate(count [26]int, digit []int, num int) [26]int {
	for i := range digit {
		count[Letters[i]-'a'] -= digit[i] * num
	}
	return count
}

func originalDigits(s string) string {
	count := [26]int{}
	for i := range s {
		count[s[i]-'a']++
	}

	digitNums := [10]int{}

	//zero count
	digitNums[0] = count['z'-'a']
	if digitNums[0] != 0 {
		count = countUpdate(count, DigitCount[0], digitNums[0])
	}

	//six count
	digitNums[6] = count['x'-'a']
	if digitNums[6] != 0 {
		count = countUpdate(count, DigitCount[6], digitNums[6])
	}
	//two count
	digitNums[2] = count['w'-'a']
	if digitNums[2] != 0 {
		count = countUpdate(count, DigitCount[2], digitNums[2])
	}
	//four count
	digitNums[4] = count['u'-'a']
	if digitNums[4] != 0 {
		count = countUpdate(count, DigitCount[4], digitNums[4])
	}
	//five count
	digitNums[5] = count['f'-'a']
	if digitNums[5] != 0 {
		count = countUpdate(count, DigitCount[5], digitNums[5])
	}
	//eight count
	digitNums[8] = count['g'-'a']
	if digitNums[8] != 0 {
		count = countUpdate(count, DigitCount[8], digitNums[8])
	}
	//seven count
	digitNums[7] = count['v'-'a']
	if digitNums[7] != 0 {
		count = countUpdate(count, DigitCount[7], digitNums[7])
	}
	//three count
	digitNums[3] = count['t'-'a']
	if digitNums[3] != 0 {
		count = countUpdate(count, DigitCount[3], digitNums[3])
	}
	//one count
	digitNums[1] = count['o'-'a']
	if digitNums[1] != 0 {
		count = countUpdate(count, DigitCount[1], digitNums[1])
	}
	//nine count
	digitNums[9] = count['e'-'a']
	if digitNums[9] != 0 {
		count = countUpdate(count, DigitCount[9], digitNums[9])
	}
	res := ""
	for i := range digitNums {
		res += strings.Repeat(string(byte(i)+'0'), digitNums[i])
	}
	return res
}

func main() {
	str := "owoztneoer"
	fmt.Println(originalDigits(str))

	str = "fviefuro"
	fmt.Println(originalDigits(str))
}
