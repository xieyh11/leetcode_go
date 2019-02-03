package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	Used     *list.List
	Keys     map[int]int
	KeysAt   map[int]*list.Element
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{list.New(), make(map[int]int), make(map[int]*list.Element), capacity}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.Keys[key]; ok {
		at := this.KeysAt[key]
		this.Used.MoveToFront(at)
		return v
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if at, ok := this.KeysAt[key]; ok {
		this.Used.MoveToFront(at)
		this.Keys[key] = value
	} else {
		if this.Used.Len() == this.Capacity {
			back := this.Used.Back()
			backKey := back.Value.(int)
			this.Used.Remove(back)
			delete(this.Keys, backKey)
			delete(this.KeysAt, backKey)
		}
		this.KeysAt[key] = this.Used.PushFront(key)
		this.Keys[key] = value
	}
}

func main() {
	cache := Constructor(2 /* capacity */)

	fmt.Println(cache.Get(2))
	cache.Put(2, 6)
	fmt.Println(cache.Get(1))
	cache.Put(1, 5)
	cache.Put(1, 2)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}
