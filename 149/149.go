package main

import (
	"../common"
	"fmt"
)

type Point = common.Point

func isOnLine(a, b, c Point) bool {
	if a.X == b.X {
		return b.X == c.X
	}
	if b.X == c.X {
		return a.X == b.X
	}
	abDeltaX := a.X - b.X
	abDeltaY := a.Y - b.Y
	bcDeltaX := b.X - c.X
	bcDeltaY := b.Y - c.Y
	return abDeltaX*bcDeltaY == abDeltaY*bcDeltaX
}

func maxPoints(points []Point) int {
	pointsNum := len(points)
	if pointsNum <= 2 {
		return pointsNum
	}
	pairs := make([][]bool, pointsNum)
	for i := range pairs {
		pairs[i] = make([]bool, pointsNum)
	}

	everyPoints := make(map[Point]int)
	noDuplicatePoints := make([]Point, 0)
	for _, point := range points {
		if _, ok := everyPoints[point]; !ok {
			noDuplicatePoints = append(noDuplicatePoints, point)
		}
		everyPoints[point]++
	}
	if len(noDuplicatePoints) <= 2 {
		return pointsNum
	}
	points = noDuplicatePoints
	pointsNum = len(points)

	maxRes := 0
	for i := range points {
		for j := i + 1; j < pointsNum; j++ {
			if !pairs[i][j] {
				onLineIdx := []int{i, j}
				for k := j + 1; k < pointsNum; k++ {
					if isOnLine(points[i], points[j], points[k]) {
						onLineIdx = append(onLineIdx, k)
					}
				}
				for m := range onLineIdx {
					for n := m + 1; n < len(onLineIdx); n++ {
						pairs[onLineIdx[m]][onLineIdx[n]] = true
					}
				}
				tempNodes := 0
				for _, idx := range onLineIdx {
					tempNodes += everyPoints[points[idx]]
				}
				if tempNodes > maxRes {
					maxRes = tempNodes
				}
			}
		}
	}
	return maxRes
}

func main() {
	points := []Point{{1, 1}, {2, 2}, {3, 3}}
	fmt.Println(maxPoints(points))
}
