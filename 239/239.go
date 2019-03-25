package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	Value *list.List
	Index *list.List
}

func NewQueue() Queue {
	return Queue{list.New(), list.New()}
}

func (a *Queue) Add(val, idx int) {
	if a.Value.Len() > 0 {
		for back := a.Value.Back().Value.(int); back <= val; back = a.Value.Back().Value.(int) {
			a.Value.Remove(a.Value.Back())
			a.Index.Remove(a.Index.Back())
			if a.Value.Len() == 0 {
				break
			}
		}
	}
	a.Value.PushBack(val)
	a.Index.PushBack(idx)
}

func (a *Queue) RemoveIndex(idx int) {
	if a.Value.Len() > 0 {
		for front := a.Index.Front().Value.(int); front <= idx; front = a.Index.Front().Value.(int) {
			a.Value.Remove(a.Value.Front())
			a.Index.Remove(a.Index.Front())
			if a.Value.Len() == 0 {
				break
			}
		}
	}
}

func (a *Queue) GetFront() (int, int) {
	return a.Value.Front().Value.(int), a.Index.Front().Value.(int)
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) < k {
		k = len(nums)
	}
	queue := NewQueue()
	for i := 0; i < k; i++ {
		queue.Add(nums[i], i)
	}
	res := make([]int, 1)
	res[0], _ = queue.GetFront()
	for i := k; i < len(nums); i++ {
		queue.Add(nums[i], i)
		queue.RemoveIndex(i - k)
		temp, _ := queue.GetFront()
		res = append(res, temp)
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
