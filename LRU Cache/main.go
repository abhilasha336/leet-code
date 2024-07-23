package main

import "fmt"

type LRUCache struct {
	capacity int
	Cache    map[int]int
	usedKeys []int
}

type Cache interface {
	Get(key int) int
	Put(key, value int)
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		Cache:    make(map[int]int),
		usedKeys: []int{},
	}
}

func (this *LRUCache) Get(key int) int {
	if value, exists := this.Cache[key]; exists {
		this.updateUsage(key)
		return value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, exists := this.Cache[key]; exists {
		this.Cache[key] = value
		this.updateUsage(key)
	} else {
		if len(this.Cache) >= this.capacity {
			// Evict the least recently used key
			lruKey := this.usedKeys[0]
			delete(this.Cache, lruKey)
			this.usedKeys = this.usedKeys[1:]
		}
		this.Cache[key] = value
		this.usedKeys = append(this.usedKeys, key)
	}
}

func (this *LRUCache) updateUsage(key int) {
	// Remove the key from the usedKeys slice
	for i, k := range this.usedKeys {
		if k == key {
			this.usedKeys = append(this.usedKeys[:i], this.usedKeys[i+1:]...)
			break
		}
	}
	// Add the key to the end to mark it as recently used
	this.usedKeys = append(this.usedKeys, key)
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)

	fmt.Println(obj.Get(1)) // returns 1
	// obj.Put(3, 3)           // evicts key 2
	// fmt.Println(obj.Get(2)) // returns -1 (not found)
	// obj.Put(4, 4)           // evicts key 1
	// fmt.Println(obj.Get(1)) // returns -1 (not found)
	// fmt.Println(obj.Get(3)) // returns 3
	// fmt.Println(obj.Get(4)) // returns 4
}
