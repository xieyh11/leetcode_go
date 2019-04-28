package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Edge struct {
	P0 Point
	P1 Point
}

func goUp(s Point, length int, border []Edge) (bool, []Edge) {
	if len(border) == 0 {
		border = append(border, Edge{s, Point{s.X, s.Y + length}})
	} else {
		if (border[1].P1.X >= s.X && border[1].P1.Y <= s.Y+length) || (border[0].P0.X == s.X && border[0].P0.Y <= s.Y+length) {
			return false, border
		} else if border[1].P1.X < s.X && border[1].P1.Y >= s.Y+length && border[0].P0.Y <= s.Y+length {
			border[2] = border[0]
		}
		border[0] = Edge{s, Point{s.X, s.Y + length}}
	}
	return true, border
}

func goLeft(s Point, length int, border []Edge) (bool, []Edge) {
	if len(border) == 1 {
		border = append(border, Edge{Point{s.X - length, s.Y}, s})
	} else {
		if (border[2].P1.Y >= s.Y && border[2].P1.X >= s.X-length) || (border[1].P1.Y == s.Y && border[1].P1.X >= s.X-length) {
			return false, border
		} else if border[2].P1.Y <= s.Y && border[2].P1.X <= s.X-length && border[1].P1.X >= s.X-length {
			border[3] = border[1]
		}
		border[1] = Edge{Point{s.X - length, s.Y}, s}
	}
	return true, border
}

func goDown(s Point, length int, border []Edge) (bool, []Edge) {
	if len(border) == 2 {
		border = append(border, Edge{Point{s.X, s.Y - length}, s})
	} else {
		if (border[3].P0.X <= s.X && border[3].P0.Y >= s.Y-length) || (border[2].P1.X == s.X && border[2].P1.Y >= s.Y-length) {
			return false, border
		} else if border[3].P0.X >= s.X && border[1].P1.Y <= s.Y-length && border[2].P1.Y >= s.Y-length {
			border[0] = border[2]
		}
		border[2] = Edge{Point{s.X, s.Y - length}, s}
	}
	return true, border
}

func goRight(s Point, length int, border []Edge) (bool, []Edge) {
	if (border[0].P0.Y <= s.Y && border[0].P0.X <= s.X+length) || (len(border) == 4 && border[3].P0.Y == s.Y && border[3].P0.X <= s.X+length) {
		return false, border
	} else if border[0].P0.Y >= s.Y && border[0].P0.X >= s.X+length {
		if len(border) == 4 && border[3].P0.X <= s.X+length {
			border[1] = border[3]
		}
	}
	if len(border) == 3 {
		border = append(border, Edge{s, Point{s.X + length, s.Y}})
	} else {
		border[3] = Edge{s, Point{s.X + length, s.Y}}
	}
	return true, border
}

func isSelfCrossing(x []int) bool {
	border := make([]Edge, 0)
	direction := 0
	nowAt := Point{0, 0}
	for i := range x {
		switch direction {
		case 0:
			avaliable, tempBorder := goUp(nowAt, x[i], border)
			if !avaliable {
				return true
			}
			border = tempBorder
			nowAt.Y += x[i]
			direction++
		case 1:
			avaliable, tempBorder := goLeft(nowAt, x[i], border)
			if !avaliable {
				return true
			}
			border = tempBorder
			nowAt.X -= x[i]
			direction++
		case 2:
			avaliable, tempBorder := goDown(nowAt, x[i], border)
			if !avaliable {
				return true
			}
			border = tempBorder
			nowAt.Y -= x[i]
			direction++
		case 3:
			avaliable, tempBorder := goRight(nowAt, x[i], border)
			if !avaliable {
				return true
			}
			border = tempBorder
			nowAt.X += x[i]
			direction = 0
		}
	}
	return false
}

func main() {
	nums0 := []int{2, 1, 1, 2}
	fmt.Println(isSelfCrossing(nums0) == true)

	nums1 := []int{1, 2, 3, 4}
	fmt.Println(isSelfCrossing(nums1) == false)

	nums2 := []int{1, 1, 1, 1}
	fmt.Println(isSelfCrossing(nums2) == true)

	nums3 := []int{1, 1, 2, 1, 1}
	fmt.Println(isSelfCrossing(nums3) == true)

	nums4 := []int{1, 1, 2, 2, 3, 1, 1}
	fmt.Println(isSelfCrossing(nums4) == true)

	nums5 := []int{1, 1, 2, 2, 3, 3, 4, 4, 10, 4, 4, 3, 3, 2, 2, 1, 1}
	fmt.Println(isSelfCrossing(nums5) == false)
}
