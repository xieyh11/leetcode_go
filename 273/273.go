package main

import (
	"fmt"
)

var lessTwenty = []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var everyTen = []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

func lessHundred(num int) string {
	if num == 0 {
		return ""
	}
	if num < 20 {
		return lessTwenty[num]
	}
	tenth := num / 10
	remain := num % 10
	res := everyTen[tenth-2]
	if remain != 0 {
		res += " " + lessTwenty[remain]
	}
	return res
}

func lessThousand(num int) string {
	res := lessHundred(num % 100)
	if num >= 100 {
		if res == "" {
			res = lessTwenty[num/100] + " Hundred"
		} else {
			res = lessTwenty[num/100] + " Hundred " + res
		}
	}
	return res
}

func numberToWords(num int) string {
	if num == 0 {
		return lessTwenty[num]
	}
	whichLevel := 0
	res := ""
	for num > 0 {
		tempR := lessThousand(num % 1000)
		if tempR != "" {
			switch whichLevel {
			case 1:
				tempR += " Thousand"
			case 2:
				tempR += " Million"
			case 3:
				tempR += " Billion"
			}
			if res == "" {
				res = tempR
			} else {
				res = tempR + " " + res
			}
		}
		num /= 1000
		whichLevel++
	}
	return res
}

func main() {
	num := 1000000
	should := "One Million"
	res := numberToWords(num)
	fmt.Println(res)
	fmt.Println(res == should)
}
