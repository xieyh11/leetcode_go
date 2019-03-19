package main

import (
	"fmt"
)

func pointInRectangle(A, B, C, D int, x, y int) bool {
	return x >= A && x <= C && y >= B && y <= D
}

func firstCoverBySecond(A int, B int, C int, D int, E int, F int, G int, H int) (int, bool) {
	secondIn := 0
	// second rectangle in first
	inPoint := make([][2]int, 0)
	if pointInRectangle(A, B, C, D, E, H) {
		secondIn++
		inPoint = append(inPoint, [2]int{E, H})
	}
	if pointInRectangle(A, B, C, D, G, H) {
		secondIn++
		inPoint = append(inPoint, [2]int{G, H})
	}
	if pointInRectangle(A, B, C, D, G, F) {
		secondIn++
		inPoint = append(inPoint, [2]int{G, F})
	}
	if pointInRectangle(A, B, C, D, E, F) {
		secondIn++
		inPoint = append(inPoint, [2]int{E, F})
	}
	switch secondIn {
	case 0:
		return 0, false
	case 1:
		if pointInRectangle(E, F, G, H, A, D) {
			inPoint = append(inPoint, [2]int{A, D})
		}
		if pointInRectangle(E, F, G, H, C, D) {
			inPoint = append(inPoint, [2]int{C, D})
		}
		if pointInRectangle(E, F, G, H, C, B) {
			inPoint = append(inPoint, [2]int{C, B})
		}
		if pointInRectangle(E, F, G, H, A, B) {
			inPoint = append(inPoint, [2]int{A, B})
		}
		if len(inPoint) != 2 {
			return 0, false
		}
		leftX, rightX := inPoint[0][0], inPoint[1][0]
		if leftX > rightX {
			leftX, rightX = rightX, leftX
		}
		bottomY, topY := inPoint[0][1], inPoint[1][1]
		if bottomY > topY {
			bottomY, topY = topY, bottomY
		}
		return (rightX - leftX) * (topY - bottomY), true
	case 2:
		// vertial
		if inPoint[0][0] == inPoint[1][0] {
			//left part
			if inPoint[0][0] == G {
				return (G - A) * (H - F), true
			}
			//right part
			if inPoint[0][0] == E {
				return (C - E) * (H - F), true
			}
		}
		//horizontal
		if inPoint[0][1] == inPoint[1][1] {
			//up part
			if inPoint[0][1] == F {
				return (G - E) * (D - F), true
			}
			//down part
			if inPoint[0][1] == H {
				return (G - E) * (H - B), true
			}
		}
		return 0, false
	case 4:
		return (G - E) * (H - F), true
	}
	return 0, false
}

func crossCover(A int, B int, C int, D int, E int, F int, G int, H int) (int, bool) {
	// first horizontal
	if B >= F && B <= H && D >= F && D <= H && E >= A && E <= C && G >= A && G <= C {
		return (D - B) * (G - E), true
	}
	//first vertical
	if A >= E && A <= G && C >= E && C <= G && F >= B && F <= D && H >= B && H <= D {
		return (C - A) * (H - F), true
	}
	return 0, false
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	sum := (C-A)*(D-B) + (G-E)*(H-F)
	area, second := firstCoverBySecond(A, B, C, D, E, F, G, H)
	if !second {
		area, second = firstCoverBySecond(E, F, G, H, A, B, C, D)
		if !second {
			area, _ = crossCover(A, B, C, D, E, F, G, H)
		}
	}
	return sum - area
}

func main() {
	point := [8]int{-5, -2, 5, 1, -3, -3, 3, 3}
	fmt.Println(computeArea(point[0], point[1], point[2], point[3], point[4], point[5], point[6], point[7]))
}
