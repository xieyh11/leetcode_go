package main

import (
	"container/list"
	"fmt"
)

type FreqEle struct {
	Freq   int
	Keys   *list.List
	KeyEle map[int]*list.Element
}

type LFUCache struct {
	KeyValue  map[int]int
	Capacity  int
	FreqQueue *list.List
	KeyEle    map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
	return LFUCache{make(map[int]int), capacity, list.New(), make(map[int]*list.Element)}
}

func (this *LFUCache) Get(key int) int {
	if v, ok := this.KeyValue[key]; ok {
		this.incKeyFreq(key)
		return v
	}
	return -1
}

func (this *LFUCache) removeLeast() {
	leastFreq := this.FreqQueue.Back()
	leastEle := leastFreq.Value.(FreqEle)
	keyRemove := leastEle.Keys.Remove(leastEle.Keys.Back()).(int)
	delete(leastEle.KeyEle, keyRemove)
	delete(this.KeyEle, keyRemove)
	delete(this.KeyValue, keyRemove)

	if leastEle.Keys.Len() == 0 {
		this.FreqQueue.Remove(leastFreq)
	} //else {
	// 	leastFreq.Value = leastEle
	// }
}

func (this *LFUCache) incKeyFreq(key int) {
	freq := this.KeyEle[key]
	freqMore := freq.Prev()
	freqEle := freq.Value.(FreqEle)

	newFreq := freqEle.Freq + 1
	freqEle.Keys.Remove(freqEle.KeyEle[key])
	delete(freqEle.KeyEle, key)

	if freqMore != nil && newFreq == freqMore.Value.(FreqEle).Freq {
		freqMoreEle := freqMore.Value.(FreqEle)
		freqMoreEle.KeyEle[key] = freqMoreEle.Keys.PushFront(key)
		// freqMore.Value = freqMoreEle
		this.KeyEle[key] = freqMore
	} else {
		keysList := list.New()
		keysEle := map[int]*list.Element{key: keysList.PushFront(key)}
		newEle := this.FreqQueue.InsertBefore(FreqEle{newFreq, keysList, keysEle}, freq)
		this.KeyEle[key] = newEle
	}

	if freqEle.Keys.Len() == 0 {
		this.FreqQueue.Remove(freq)
	} // else {
	// 	freq.Value = freqEle
	// }
}

func (this *LFUCache) insertNewKey(key int) {
	leastFreq := this.FreqQueue.Back()

	if leastFreq != nil && leastFreq.Value.(FreqEle).Freq == 1 {
		leastEle := leastFreq.Value.(FreqEle)
		leastEle.KeyEle[key] = leastEle.Keys.PushFront(key)
		// leastFreq.Value = leastEle
		this.KeyEle[key] = leastFreq
	} else {
		keyList := list.New()
		keyEle := map[int]*list.Element{key: keyList.PushFront(key)}
		newEle := this.FreqQueue.PushBack(FreqEle{1, keyList, keyEle})
		this.KeyEle[key] = newEle
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.Capacity == 0 {
		return
	}
	if _, ok := this.KeyValue[key]; ok {
		this.KeyValue[key] = value
		this.incKeyFreq(key)
	} else {
		if this.Capacity == len(this.KeyValue) {
			this.removeLeast()
		}
		this.KeyValue[key] = value
		this.insertNewKey(key)
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3) // evicts key 2
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	cache.Put(4, 4) // evicts key 1.
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))

	cache1 := Constructor(1)
	cache1.Put(2, 1)
	fmt.Println(cache1.Get(2))
	cache1.Put(3, 2) // evicts key 2
	fmt.Println(cache1.Get(2))
	fmt.Println(cache1.Get(3))
}
