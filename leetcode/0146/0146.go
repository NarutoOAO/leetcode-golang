package _146

import "container/list"

type LRUCache struct {
	cap int
	m   map[int]*list.Element
	l   *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		m:   make(map[int]*list.Element),
		l:   list.New(),
	}
}

type pair struct {
	k, v int
}

func (this *LRUCache) Get(key int) int {
	if element, ok := this.m[key]; ok {
		this.l.MoveToFront(element)
		return this.m[key].Value.(pair).v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if element, ok := this.m[key]; ok {
		element.Value = pair{key, value}
		this.l.MoveToFront(element)
		this.m[key] = element
	} else {
		element := this.l.PushFront(pair{key, value})
		this.m[key] = element
	}
	if this.l.Len() > this.cap {
		element := this.l.Back()
		this.l.Remove(element)
		delete(this.m, element.Value.(pair).k)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
