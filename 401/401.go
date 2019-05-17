package main

import (
	"fmt"
)

var Hours = [][]string{{"0"}, {"1", "2", "4", "8"}, {"3", "5", "6", "9", "10"}, {"7", "11"}}
var Minutes = [][]string{{"00"}, {"01", "02", "04", "08", "16", "32"}, {"03", "05", "06", "09", "10", "12", "17", "18", "20", "24", "33", "34", "36", "40", "48"}, {"07", "11", "13", "14", "19", "21", "22", "25", "26", "28", "35", "37", "38", "41", "42", "44", "49", "50", "52", "56"}, {"15", "23", "27", "29", "30", "39", "43", "45", "46", "51", "53", "54", "57", "58"}, {"31", "47", "55", "59"}}

func readBinaryWatch(num int) []string {
	res := []string{}
	for i := 0; i < len(Hours) && i <= num; i++ {
		j := num - i
		if j < len(Minutes) {
			for _, h := range Hours[i] {
				for _, m := range Minutes[j] {
					res = append(res, h+":"+m)
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(readBinaryWatch(4))
}
