package main

import (
	"container/list"
	"fmt"
)

type MidPointer struct {
	Num int
	Idx int
	Ele *list.Element
}

// 0 equal, 1 larger than num, -1 samller than num
func (this *MidPointer) Compare(num int) int {
	if this.Ele != nil {
		midValue := this.Ele.Value.(int)
		if midValue > num {
			return 1
		} else if midValue < num {
			return -1
		} else {
			return 0
		}
	}
	if this.Num > num {
		return 1
	} else if this.Num < num {
		return -1
	} else {
		return 0
	}
}

func (this *MidPointer) Value() float64 {
	if this.Ele != nil {
		return float64(this.Ele.Value.(int))
	}
	return float64(this.Num)
}

type MedianFinder struct {
	ZeroToHundred [101]int
	Larger        *list.List
	Smaller       *list.List
	Mid1          MidPointer
	Mid2          MidPointer
	Len           int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{Larger: list.New(), Smaller: list.New()}
}

func (this *MedianFinder) MoveRight(mid MidPointer) MidPointer {
	if mid.Ele != nil {
		mid.Ele = mid.Ele.Next()
		if mid.Ele == nil {
			found := false
			for i := range this.ZeroToHundred {
				if this.ZeroToHundred[i] != 0 {
					mid = MidPointer{Num: i, Idx: 0}
					found = true
					break
				}
			}
			if !found {
				mid = MidPointer{Ele: this.Larger.Front()}
			}
		}
		return mid
	}
	mid.Idx++
	if this.ZeroToHundred[mid.Num] == mid.Idx {
		found := false
		for i := mid.Num + 1; i <= 100; i++ {
			if this.ZeroToHundred[i] != 0 {
				mid = MidPointer{Num: i, Idx: 0}
				found = true
				break
			}
		}
		if !found {
			mid = MidPointer{Ele: this.Larger.Front()}
		}
	}
	return mid
}

func (this *MedianFinder) MoveLeft(mid MidPointer) MidPointer {
	if mid.Ele != nil {
		mid.Ele = mid.Ele.Prev()
		if mid.Ele == nil {
			found := false
			for i := 100; i >= 0; i-- {
				if this.ZeroToHundred[i] != 0 {
					mid = MidPointer{Num: i, Idx: this.ZeroToHundred[i] - 1}
					found = true
					break
				}
			}
			if !found {
				mid = MidPointer{Ele: this.Smaller.Back()}
			}
		}
		return mid
	}
	mid.Idx--
	if mid.Idx < 0 {
		found := false
		for i := mid.Num - 1; i >= 0; i-- {
			if this.ZeroToHundred[i] != 0 {
				mid = MidPointer{Num: i, Idx: this.ZeroToHundred[i] - 1}
				found = true
				break
			}
		}
		if !found {
			mid = MidPointer{Ele: this.Smaller.Back()}
		}
	}
	return mid
}

func (this *MedianFinder) AddLargerOne() {
	if this.Mid1 == this.Mid2 {
		this.Mid2 = this.MoveRight(this.Mid2)
	} else {
		this.Mid1 = this.Mid2
	}
}

func (this *MedianFinder) AddSmallerOne() {
	if this.Mid1 == this.Mid2 {
		this.Mid1 = this.MoveLeft(this.Mid1)
	} else {
		this.Mid2 = this.Mid1
	}
}

func insertBefore(lt *list.List, ele *list.Element, num int) {
	if ele == nil {
		ele = lt.Back()
		if ele == nil || ele.Value.(int) <= num {
			lt.PushBack(num)
			return
		}
	}
	for ele.Prev() != nil && ele.Prev().Value.(int) > num {
		ele = ele.Prev()
	}
	lt.InsertBefore(num, ele)

}

func insertAfter(lt *list.List, ele *list.Element, num int) {
	if ele == nil {
		ele = lt.Front()
		if ele == nil || ele.Value.(int) >= num {
			lt.PushFront(num)
			return
		}
	}
	for ele.Next() != nil && ele.Next().Value.(int) < num {
		ele = ele.Next()
	}
	lt.InsertAfter(num, ele)
}

func (this *MedianFinder) AddNum(num int) {
	if this.Len == 0 {
		if num < 0 {
			this.Smaller.PushBack(num)
			this.Mid1.Ele = this.Smaller.Front()
			this.Mid2.Ele = this.Smaller.Front()
		} else if num > 100 {
			this.Larger.PushBack(num)
			this.Mid1.Ele = this.Larger.Front()
			this.Mid2.Ele = this.Larger.Front()
		} else {
			this.ZeroToHundred[num]++
			this.Mid1.Num = num
			this.Mid2.Num = num
		}
		this.Len++
		return
	}
	mid1Compare := this.Mid1.Compare(num)
	mid2Compare := this.Mid2.Compare(num)
	if mid1Compare >= 0 {
		if num < 0 {
			if this.Mid1.Compare(0) < 0 {
				insertBefore(this.Smaller, this.Mid1.Ele, num)
			} else {
				insertBefore(this.Smaller, nil, num)
			}
		} else if num <= 100 {
			this.ZeroToHundred[num]++
			if mid1Compare == 0 {
				this.Mid1.Idx++
			}
			if mid2Compare == 0 {
				this.Mid2.Idx++
			}
		} else {
			insertBefore(this.Larger, this.Mid1.Ele, num)
		}
		this.AddSmallerOne()
	} else if mid2Compare <= 0 {
		if num > 100 {
			if this.Mid2.Compare(100) > 0 {
				insertAfter(this.Larger, this.Mid2.Ele, num)
			} else {
				insertAfter(this.Larger, nil, num)
			}
		} else if num >= 0 {
			this.ZeroToHundred[num]++
		} else {
			insertAfter(this.Smaller, this.Mid2.Ele, num)
		}
		this.AddLargerOne()
	} else {
		if num < 0 {
			this.Mid1 = MidPointer{Ele: this.Smaller.InsertAfter(num, this.Mid1.Ele)}
		} else if num <= 100 {
			this.ZeroToHundred[num]++
			this.Mid1 = MidPointer{Num: num, Idx: 0}
		} else {
			this.Mid1 = MidPointer{Ele: this.Larger.InsertBefore(num, this.Mid2.Ele)}
		}
		this.Mid2 = this.Mid1
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Mid1 == this.Mid2 {
		return this.Mid1.Value()
	}
	return (this.Mid1.Value() + this.Mid2.Value()) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	median := Constructor()
	nums := []int{12, 10, 13, 11, 5, 15, 1, 11, 6, 17, 14, 8, 17, 6, 4, 16, 8, 10, 2, 12, 0}
	for _, num := range nums {
		median.AddNum(num)
		fmt.Println(median.FindMedian())
	}
}
