package main

import (
	"container/list"
	"fmt"
)

type KeyPair struct {
	Key string
	Val int
}

type AllOne struct {
	Keys     []*list.List
	KeyMap   map[string]*list.Element
	LevelIn  *list.List
	LevelMap map[int]*list.Element
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{KeyMap: make(map[string]*list.Element), LevelIn: list.New(), LevelMap: make(map[int]*list.Element)}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if v, ok := this.KeyMap[key]; ok {
		pair := v.Value.(KeyPair)
		if pair.Val+1 >= len(this.Keys) {
			this.Keys = append(this.Keys, list.New())
		}
		this.Keys[pair.Val].Remove(v)
		this.KeyMap[key] = this.Keys[pair.Val+1].PushBack(KeyPair{pair.Key, pair.Val + 1})
		level := this.LevelMap[pair.Val]
		if this.Keys[pair.Val+1].Len() == 1 {
			this.LevelMap[pair.Val+1] = this.LevelIn.InsertAfter(pair.Val+1, level)
		}
		if this.Keys[pair.Val].Len() == 0 {
			this.LevelIn.Remove(level)
			delete(this.LevelMap, pair.Val)
		}
	} else {
		if len(this.Keys) == 0 {
			this.Keys = append(this.Keys, list.New())
		}
		v = this.Keys[0].PushBack(KeyPair{key, 0})
		this.KeyMap[key] = v
		if this.Keys[0].Len() == 1 {
			this.LevelMap[0] = this.LevelIn.PushFront(0)
		}
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if v, ok := this.KeyMap[key]; ok {
		pair := v.Value.(KeyPair)
		this.Keys[pair.Val].Remove(v)
		level := this.LevelMap[pair.Val]
		if pair.Val > 0 {
			this.KeyMap[key] = this.Keys[pair.Val-1].PushBack(KeyPair{pair.Key, pair.Val - 1})
			if this.Keys[pair.Val-1].Len() == 1 {
				this.LevelMap[pair.Val-1] = this.LevelIn.InsertBefore(pair.Val-1, level)
			}
		} else {
			delete(this.KeyMap, key)
		}
		if this.Keys[pair.Val].Len() == 0 {
			this.LevelIn.Remove(level)
			delete(this.LevelMap, pair.Val)
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.LevelIn.Len() == 0 {
		return ""
	}
	levelMax := this.LevelIn.Back().Value.(int)
	return this.Keys[levelMax].Front().Value.(KeyPair).Key
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.LevelIn.Len() == 0 {
		return ""
	}
	levelMin := this.LevelIn.Front().Value.(int)
	return this.Keys[levelMin].Front().Value.(KeyPair).Key
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

func main() {
	sol := Constructor()
	sol.Inc("hello")
	sol.Inc("hello")
	fmt.Println(sol.GetMaxKey())
	fmt.Println(sol.GetMinKey())

	sol.Inc("leet")
	fmt.Println(sol.GetMaxKey())
	fmt.Println(sol.GetMinKey())

	sol.Dec("leet")
	fmt.Println(sol.GetMaxKey())
	fmt.Println(sol.GetMinKey())
}
