package main

import (
	"fmt"
	"sort"
)

type Sortbuilds [][]int

func (this Sortbuilds) Len() int      { return len(this) }
func (this Sortbuilds) Swap(i, j int) { this[i], this[j] = this[j], this[i] }
func (this Sortbuilds) Less(i, j int) bool {
	if this[i][2] == this[j][2] {
		return this[i][0] < this[j][0]
	}
	return this[i][2] > this[j][2]
}

func findInRegion(region [][]int, left int) (int, bool) {
	if len(region) == 1 {
		if region[0][0] <= left && region[0][1] >= left {
			return 0, true
		} else if region[0][0] > left {
			return -1, false
		} else {
			return 0, false
		}
	}
	if len(region) == 2 {
		switch {
		case region[0][0] > left:
			return -1, false
		case region[0][0] <= left && region[0][1] >= left:
			return 0, true
		case region[0][1] < left && region[1][0] > left:
			return 0, false
		case region[1][0] <= left && region[1][1] >= left:
			return 1, true
		default:
			return 1, false
		}
	}
	mid := len(region) / 2
	if region[mid][0] <= left && region[mid][1] >= left {
		return mid, true
	} else if region[mid][0] > left {
		return findInRegion(region[:mid], left)
	} else {
		idx, isin := findInRegion(region[mid+1:], left)
		return idx + mid + 1, isin
	}
}

func splitRegion(region [][]int, building []int) [][]int {
	leftIdx, leftIsin := findInRegion(region, building[0])
	rightIdx, rightIsin := findInRegion(region, building[1])

	if leftIsin {
		for leftIdx < rightIdx-1 && region[leftIdx][1] == region[leftIdx+1][0] {
			leftIdx++
		}
		if leftIdx == rightIdx-1 && !rightIsin && region[leftIdx][1] == region[leftIdx+1][0] {
			leftIdx++
		}
	}
	if rightIsin {
		for rightIdx > leftIdx+1 && region[rightIdx][0] == region[rightIdx-1][1] {
			rightIdx--
		}
	}

	split := make([][]int, 0)
	formerLeft := building[0]
	if leftIsin {
		formerLeft = region[leftIdx][1]
	}
	for i := leftIdx + 1; i < rightIdx; i++ {
		if formerLeft < region[i][0] {
			split = append(split, []int{formerLeft, region[i][0], building[2]})

		}
		split = append(split, region[i])
		formerLeft = region[i][1]
	}
	if (rightIsin && formerLeft < region[rightIdx][0]) || (rightIdx >= 0 && formerLeft < region[rightIdx][0]) {
		split = append(split, []int{formerLeft, region[rightIdx][0], building[2]})
		formerLeft = region[rightIdx][1]
	}
	if rightIsin {
		split = append(split, region[rightIdx:]...)
	} else {
		if rightIdx >= 0 {
			split = append(split, region[rightIdx])
			if formerLeft < region[rightIdx][1] {
				formerLeft = region[rightIdx][1]
			}
		}
		split = append(split, []int{formerLeft, building[1], building[2]})
		split = append(split, region[rightIdx+1:]...)
	}
	if rightIdx != leftIdx {
		region = append(region[:leftIdx+1], split...)
	} else {
		if leftIdx >= 0 {
			region = append(region[:leftIdx], split...)
		} else {
			region = split
		}
	}
	return region
}

func mergeBuildings(buildings [][]int) [][]int {
	res := make([][]int, 0)
	i := 0
	for i < len(buildings) {
		j := i + 1
		maxRight := buildings[i][1]
		for j < len(buildings) {
			if buildings[j][2] != buildings[j-1][2] || buildings[j][0] > buildings[j-1][1] {
				break
			}
			if maxRight < buildings[j][1] {
				maxRight = buildings[j][1]
			}
			j++
		}
		if j > i+1 {
			res = append(res, []int{buildings[i][0], maxRight, buildings[i][2]})
		} else {
			res = append(res, buildings[i])
		}
		i = j
	}
	return res
}

func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}
	sort.Sort(Sortbuilds(buildings))
	buildings = mergeBuildings(buildings)
	if len(buildings) == 1 {
		return [][]int{{buildings[0][0], buildings[0][2]}, {buildings[0][1], 0}}
	}
	region := make([][]int, 1)
	region[0] = buildings[0]
	for i := 1; i < len(buildings); i++ {
		region = splitRegion(region, buildings[i])
	}
	res := make([][]int, 0)
	for i := range region {
		res = append(res, []int{region[i][0], region[i][2]})
		if i == len(region)-1 || region[i][1] < region[i+1][0] {
			res = append(res, []int{region[i][1], 0})
		}
	}
	return res
}

func main() {
	buildings := [][]int{{2190, 661048, 758784}, {9349, 881233, 563276}, {12407, 630134, 38165}, {22681, 726659, 565517}, {31035, 590482, 658874}, {41079, 901797, 183267}, {41966, 103105, 797412}, {55007, 801603, 612368}, {58392, 212820, 555654}, {72911, 127030, 629492}, {73343, 141788, 686181}, {83528, 142436, 240383}, {84774, 599155, 787928}, {106461, 451255, 856478}, {108312, 994654, 727797}, {126206, 273044, 692346}, {134022, 376405, 472351}, {151396, 993568, 856873}, {171466, 493683, 664744}, {173068, 901140, 333376}, {179498, 667787, 518146}, {182589, 973265, 394689}, {201756, 900649, 31050}, {215635, 818704, 576840}, {223320, 282070, 850252}, {252616, 974496, 951489}, {255654, 640881, 682979}, {287063, 366075, 76163}, {291126, 900088, 410078}, {296928, 373424, 41902}, {297159, 357827, 174187}, {306338, 779164, 565403}, {317547, 979039, 172892}, {323095, 698297, 566611}, {323195, 622777, 514005}, {333003, 335175, 868871}, {334996, 734946, 720348}, {344417, 952196, 903592}, {348009, 977242, 277615}, {351747, 930487, 256666}, {363240, 475567, 699704}, {365620, 755687, 901569}, {369650, 650840, 983693}, {370927, 621325, 640913}, {371945, 419564, 330008}, {415109, 890558, 606676}, {427304, 782478, 822160}, {439482, 509273, 627966}, {443909, 914404, 117924}, {446741, 853899, 285878}, {480389, 658623, 986748}, {545123, 873277, 431801}, {552469, 730722, 574235}, {556895, 568292, 527243}, {568368, 728429, 197654}, {593412, 760850, 165709}, {598238, 706529, 500991}, {604335, 921904, 990205}, {627682, 871424, 393992}, {630315, 802375, 714014}, {657552, 782736, 175905}, {701356, 827700, 70697}, {712097, 737087, 157624}, {716678, 889964, 161559}, {724790, 945554, 283638}, {761604, 840538, 536707}, {776181, 932102, 773239}, {855055, 983324, 880344}}
	fmt.Println(getSkyline(buildings))
}
