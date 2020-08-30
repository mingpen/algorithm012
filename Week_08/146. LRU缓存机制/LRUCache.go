package main

import "container/list"

//https://leetcode-cn.com/problems/lru-cache/submissions/

func main() {

}

// LRUCache ...
type LRUCache struct {
	l        *list.List
	m        map[int]*list.Element
	capacity int
}
type v struct {
	key   int
	value int
}

// Constructor  ...
func Constructor(capacity int) LRUCache {
	l := list.New()
	return LRUCache{l: l, capacity: capacity, m: make(map[int]*list.Element, capacity)}
}

// Get  ...
func (this *LRUCache) Get(key int) int {
	if e, ok := this.m[key]; ok {
		this.l.MoveToFront(e)
		return e.Value.(*v).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	e, found := this.m[key]
	defer func() {
		this.m[key] = e
	}()
	if !found && this.l.Len() < this.capacity {
		e = this.l.PushFront(&v{key: key, value: value})
		return
	}
	if !found {
		e = this.l.Back()
		delete(this.m, e.Value.(*v).key)
		e.Value.(*v).key = key
	}
	e.Value.(*v).value = value
	this.l.MoveToFront(e)
	// found
	// not found
	// if len+1 < cap
	// inster
	// else
	// mv and set
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
